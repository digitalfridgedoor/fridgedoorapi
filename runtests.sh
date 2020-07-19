allExitcodes=()

go test .
allExitcodes+=$?

go test ./fridgedoorgateway
allExitcodes+=$?

go test ./clippingapi
allExitcodes+=$?

go test ./imageapi
allExitcodes+=$?

go test ./linkeduserapi
allExitcodes+=$?

go test ./planapi
allExitcodes+=$?

go test ./recipeapi
allExitcodes+=$?

go test ./userviewapi
allExitcodes+=$?

for t in ${allExitcodes[@]}; do
  if [[ $t != 0 ]]
    then exit $t
  fi
done
