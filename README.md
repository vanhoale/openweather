# openweather
a openweather proxy api

# Run code locally:
go run main.go

# Deploy to Kubernetes
You should install helm and kind at your local machine

cd $PROJECT_HOME/charts/openweather

helm upgrade --install openweather -n $APP_NAMESPACE .
