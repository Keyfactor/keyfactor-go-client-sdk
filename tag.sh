RC_VERSION=rc.13
#TAG_VERSION_10=v10.0.0-$RC_VERSION
#TAG_VERSION_11=v11.0.0-$RC_VERSION
TAG_VERSION_1=v1.1.0-$RC_VERSION
#git tag -d $TAG_VERSION_10 || true
#git tag -d $TAG_VERSION_11 || true
git tag -d $TAG_VERSION_1 || true
#git push personal :$TAG_VERSION
#git tag $TAG_VERSION_10
#git push origin $TAG_VERSION_10
#git tag $TAG_VERSION_11
#git push origin $TAG_VERSION_11
git tag $TAG_VERSION_1
git push origin $TAG_VERSION_1