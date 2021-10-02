if [ -z "$(git status --porcelain)" ]; then 
  # Working directory clean
  echo 'Working directory is clean, starting to deploy'
else
  # Uncommitted changes
  echo
  echo "!!!!!!!!! ERROR !!!!!!!!!"
  echo "Failed to deploy server code! There are uncommitted changes."
  echo "Commit or discard the changes and try again!"
  exit 1
fi

./get-logs.sh

git push deploy-server main
