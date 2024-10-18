# Welcome to Kasa API

## Quickstart

1. Create a .env file at the root of your project and complete it with the fields indicated in the .env.dist file.

2. Migrate the database with the following command :
    ```bash
    go run github.com/steebchen/prisma-client-go db push
    ```
3. Use docker-compose to launch the API:
    ```bash
    docker compose up # To just launch the API.
    # or
    docker compose watch # To launch the API with the watcher.
    ```

