set -e

./scripts/apply_env.sh ./src/packages/varsoapi/package.json
cd src/packages/varsoapi
npm config set registry https://npm.pkg.github.com/
npm unpublish @luminancetech/varso@${GIT_REF} || echo No version to unpublish
npm publish