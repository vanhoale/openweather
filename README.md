# openweather
a openweather proxy api

# Run code locally:
go run main.go

export OPENWEATHER_API_KEY=$YOUR_KEY && curl -H 'Content-Type: application/json'  "http://localhost:8080/weather?lat=44.9706674&lon=-93.3438792"

# Deploy to Kubernetes
You should install helm and kind at your local machine

cd $PROJECT_HOME/charts/openweather

helm upgrade --install openweather -n $APP_NAMESPACE --set openweatherApiKey=$YOUR_KEY .
