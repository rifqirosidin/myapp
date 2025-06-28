echo "📦 Building docker image"
docker build -t myapp-go:ci .

echo "🚀 Re-deploying with docker-compose..."
docker-compose --env-file .env.production up -d --build

echo "✅ Done. Access at: https://myapp.local"