#!/bin/bash

# ================== START: Inputs ===================================
SWAGGER_API_BASE="https://matthew-irby-4.kftestlab.com/Keyfactor/API/swagger" # The base URL for the Command's Swagger docs
COMMAND_VERSION_NAME="v25" # This will be the folder name for the generated SDK, and should match the version of Keyfactor Command you are generating the SDK for.
COMMAND_VERSION_FULL="25.1.1" # This can be found on the Command website and also via the /License endpoint (under 'KeyfactorVersion')
API_VERSIONS=("1" "2") # Keyfactor API Versions you will want to generate an SDK for
# ================== END: Inputs ===================================


# check if openapi-generator is installed
if ! command -v openapi-generator &> /dev/null; then
  echo "openapi-generator could not be found. Please install it first."
  exit 1
fi

openapi_generator_version=$(openapi-generator version)
current_year=$(date +%Y)
config_file_name="openapi-config.yml"
generator_language="go"
custom_templates_path="custom-templates/go"
base_dir=$(pwd)

# Remove trailing "/" character
command_api_swagger_base_path=${SWAGGER_API_BASE%/}

echo "Deleting existing folder path $base_dir/$COMMAND_VERSION_NAME" <&2
rm -rf ./${COMMAND_VERSION_NAME}

# Dynamically generated variables populated by API client version
imports=""
clientInstatantiation=""
clientList=""
documentationList=""
apiClientExamples=""

for version_num in "${API_VERSIONS[@]}"
do
  # Make sure we are back in the original directory
  cd $base_dir

  version="v$version_num"

  api_version_path="${COMMAND_VERSION_NAME}/api/keyfactor/${version}"
  echo "Creating path $base_dir/$api_version_path" <&2
  mkdir -p ./$api_version_path

  request_url="$command_api_swagger_base_path/$version/swagger.json"
  spec_file="Keyfactor-Command-${COMMAND_VERSION_NAME}-${COMMAND_VERSION_FULL}-spec_$version.json"
  spec_raw_file="Keyfactor-Command-${COMMAND_VERSION_NAME}-${COMMAND_VERSION_FULL}-spec_$version-raw.json"

  echo "Getting Swagger JSON at $request_url"

  http_code=$(curl -s -X GET $request_url -o $spec_file -w "%{http_code}")

  if [ $http_code != 200 ]; then
    echo "ERROR: Received an HTTP status of $http_code from $request_url. Exiting."

    # TODO: Run cleanup

    exit 1
  fi

  echo "Swagger JSON was retrieved successfully." <&2

  echo "Backing up original spec file..." <&2
  cp $spec_file $spec_raw_file
  echo "Done." <&2

  echo "Modifying Swagger doc with default schema type on post parameters..." <&2

  # There are some parameters missing a schema, which causes a failure in the generator. 
  # If missing, then force it to have a schema and assume it's a string.
  jq '
    .paths |= with_entries(
      if .value.post and .value.post.parameters then
        .value.post.parameters |= map(
          if has("schema") then . else . + { "schema": { "type": "string" } } end
        )
      else
        .
      end
    )
  ' $spec_file > modified_$spec_file

  echo "Done." <&2

  cp modified_$spec_file $spec_file

  echo "Modifying Swagger doc with updated server url..." <&2

  # Replace the Keyfactor Command URL used with the example URL
  jq '
    .servers |= map(
      .url = "http://keyfactor.example.com"
    )
  ' $spec_file > modified_$spec_file

  echo "Done." <&2

  cp modified_$spec_file $spec_file

  echo "Modifying Swagger doc with updated version..." <&2

  jq --arg version "$(printf "$version_num")" \
    '.info.version = $version' $spec_file > modified_$spec_file

  cp modified_$spec_file $spec_file

  echo "Done." <&2

  echo "Updating operationId on paths..." <&2

  # For each HTTP method, for clearer intent replace the method name with its purpose
  # If the path contains an "id" parameter, append "ById" to the operationId
  jq --arg version "$(printf "$version_num")" \
  '
  # Iterate through all paths
  .paths |= with_entries(
    . as $path |
    .value |= with_entries(
        .value.operationId = (
          # (.value.tags[0] // "DefaultTag") + "_" +
            (if .key == "post" then "Create" 
             elif .key == "put" then "Update" 
             elif .key == "delete" then "Delete" 
             elif .key == "patch" then "Patch" 
             else "Get" end) +
          (
            $path.key | gsub("[{}]"; "") | split("/") | map(select(. != "")) | map(
              if . == "id" then "ById" 
              else . 
              end
            ) | join("_")
          ) 
          # + "V" + $version
        )
    )
  )' $spec_file > modified_$spec_file

  cp modified_$spec_file $spec_file

  echo "Done." <&2

  echo "Fixing the namespace declaration on models..." <&2

  # Replace superfluous model names to not include the model namespace
  sed 's/Keyfactor\.Web\.KeyfactorApi\.Models\.//g' $spec_file > modified_$spec_file

  cp modified_$spec_file $spec_file

  echo "Done." <&2

  echo "Updating enum values..." <&2

  # x-enum-varnames makes it easy to map the enum number to a string value in the generated code.
  # Most enums in the API spec are missing this, so we need to add it manually.
  jq '
    .components.schemas."CSS.CMS.Core.Enums.ClaimType" |=
     if has("x-enum-varnames") then .
     else (. + {"x-enum-varnames": ["User", "Group", "Computer", "OAuthOid", "OAuthRole", "OAuthSubject", "OAuthClientId" ] })
     end
  ' $spec_file > modified_$spec_file

  cp modified_$spec_file $spec_file

  jq '
    .components.schemas."System.DayOfWeek" |= 
     if has("x-enum-varnames") then .
     else (. + {"x-enum-varnames": ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" ] })
     end
  ' $spec_file > modified_$spec_file

  cp modified_$spec_file $spec_file

  # Hotfix -- update the enums array of enrollment types. Zendesk tix 139784
  jq 'if .components.schemas["CSS.CMS.Core.Enums.EnrollmentType"] 
    then .components.schemas["CSS.CMS.Core.Enums.EnrollmentType"].enum = [0, 1, 2, 3, 4, 5, 6, 7] 
    else . end' $spec_file > modified_$spec_file

  cp modified_$spec_file $spec_file

  echo "Done." <&2

  rm modified_$spec_file

  echo "Running openapi-generator..." <&2

  if [[ ! $(openapi-generator generate \
    -g $generator_language \
    -i $spec_file \
    -t $custom_templates_path \
    -c $config_file_name \
    -o $api_version_path \
    -p sdkVersion=$COMMAND_VERSION_NAME \
    -p apiVersion=$COMMAND_VERSION_FULL \
    -p apiNameSuffix="Api" \
    -p packageName="${version}" \
    -p openApiVersion=$openapi_generator_version \
    -p currentYear=$current_year)  ]]; then
    echo "Generator failed to complete. Exiting."

    # TODO: Cleanup

    exit 1
  fi

  echo "Generator completed successfully." <&2

  cd $api_version_path

  echo "Go dependencies successfully installed."

  echo "Tidying up imports..."

  goimports -w .
  go mod tidy

  echo "Removing test directory..." <&2

  rm -rf ./test

  echo "Done." <&2

  echo "Copying config.yml to the directory..." <&2

  cp $base_dir/config.yml .

  echo "Done." <&2

  imports+="$version \"github.com\/Keyfactor\/keyfactor-go-client-sdk\/$COMMAND_VERSION_NAME\/api\/keyfactor\/${version}\"\n"
  clientInstantiation+="clientV$version_num, err := $version.NewAPIClient(cfg)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n"
  clientList+="\tV$version_num: clientV$version_num,\n"
  documentationList+="- [$version API Documentation](.\/api\/keyfactor\/$version\/README.md)\n"
  apiClientExamples+="\/\/ api := client.V$version_num.ExampleApi \/\/ Access V$version_num API Service\n\/\/ req := api.NewUpdateExampleRequest(ctx) \/\/ Build new API request\n\/\/ req = req.ExampleUpdateRequest(V$version_num.ExampleUpdateRequest{ Name: \"Hello\" }) \/\/ Add a body to the API request (if applicable)\n\/\/ resp, httpResp, err := req.Execute() \/\/ Execute request\n\n"

  echo "Deleting unnecessary files..." <&2
  rm go.mod go.sum .travis.yml
done

echo "OpenAPI SDKS generated successfully. Post-processing..." <&2

cd $base_dir/$COMMAND_VERSION_NAME

echo "Creating go.mod and go.sum files..." <&2
cp $base_dir/custom-templates/postprocess/go/go.mod.template go.mod
cp $base_dir/custom-templates/postprocess/go/go.sum.template go.sum


echo "Creating base client.go file..." <&2
cp $base_dir/custom-templates/postprocess/go/client.go.template client.go

echo "Creating base README.md file..." <&2
cp $base_dir/custom-templates/postprocess/go/README.md.template README.md

echo "Creating base helpers.go file..." <&2
cp $base_dir/custom-templates/postprocess/go/helpers.go.template helpers.go

echo "Updating packageName in go.mod..." <&2
sed "s/{{packageVersion}}/$COMMAND_VERSION_NAME/g" go.mod > go.mod.tmp
mv go.mod.tmp go.mod

# === client.go updates =====

echo "Updating imports in client.go..." <&2
sed "s/{{imports}}/$imports/g" client.go > client.go.tmp
mv client.go.tmp client.go

echo "Updating client instantiation in client.go..." <&2
sed "s/{{clientInstantiation}}/$clientInstantiation/g" client.go > client.go.tmp
mv client.go.tmp client.go

echo "Updating client list in client.go..." <&2
sed "s/{{clientList}}/$clientList/g" client.go > client.go.tmp
mv client.go.tmp client.go

# =============================

# === helpers.go updates ======


echo "Updating packageVersion in helpers.go..." <&2
sed "s/{{packageVersion}}/$COMMAND_VERSION_NAME/g" helpers.go > helpers.go.tmp
mv helpers.go.tmp helpers.go

# =============================


echo "Updating packageName in README.md..." <&2
sed "s/{{packageVersion}}/$COMMAND_VERSION_NAME/g" README.md > README.md.tmp
mv README.md.tmp README.md

echo "Updating version in README.md..." <&2
sed "s/{{version}}/$COMMAND_VERSION_FULL/g" README.md > README.md.tmp
mv README.md.tmp README.md

echo "Updating documentationList in README.md..." <&2
sed "s/{{documentationList}}/$documentationList/g" README.md > README.md.tmp
mv README.md.tmp README.md

echo "Updating apiClientExamples in README.md..." <&2
sed "s/{{apiClientExamples}}/$apiClientExamples/g" README.md > README.md.tmp
mv README.md.tmp README.md

echo "Running go mod tidy..."
go mod tidy

echo "Running go fmt..." <&2
go fmt .

echo "Done post-processing." <&2

echo "Script has completed successfully!" <&2
