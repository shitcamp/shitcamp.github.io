doctl apps logs ${APP_ID} shitcamp-api --type run > ./logs/shitcamp-latest.log

START_TIME=$(grep "[GIN]" logs/shitcamp-latest.log | head -n1 | cut -d' ' -f2 | cut -d. -f1)
END_TIME=$(grep "[GIN]" logs/shitcamp-latest.log | tail -n1 | cut -d' ' -f2 | cut -d. -f1)

LOG_NAME=$(echo "shitcamp-${START_TIME}_to_${END_TIME}.log")
LOG_NAME=$(echo $(sed 's/:/-/g' <<<"$LOG_NAME"))

mv ./logs/shitcamp-latest.log ./logs/${LOG_NAME}
