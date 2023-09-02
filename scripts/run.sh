buildCmd="docker-compose"
echo "Building docker containers..."
echo "Current directory: $"

for f in /build/docker/*.yaml; do
  buildCmd="$buildCmd -f $f"
done

for f in /build/docker/*.yaml; do
  buildCmd="$buildCmd -f $f"
done
bash "$buildCmd up"