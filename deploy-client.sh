# if [ -z "$(git status --untracked-files=no --porcelain)" ]; then # Working directory clean excluding untracked files
if [ -z "$(git status --porcelain)" ]; then 
  # Working directory clean
  echo 'Working directory is clean, starting to deploy'
else
  # Uncommitted changes
  echo
  echo "!!!!!!!!! ERROR !!!!!!!!!"
  echo "Failed to deploy client code!! There are uncommitted changes."
  echo "Commit or discard the changes and try again!"
  exit 1
fi

cd web
npm version patch
git reset
git add package.json package-lock.json
git commit -m 'Update npm app version'

npm install
npm run deploy
