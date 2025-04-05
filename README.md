# Titly

Titly is a blazingly fast URL shortner written in go.

## Tech Stack

- Go
- Redis
- Sveltekit

## Run the Backend without Docker

1. Ensure you have Go installed on your system.

2. Run the application:

  ```bash
  go run .
  ```

3. The backend will be available at `http://localhost:4000`.

## Run the Backend using Docker

1. Build the Docker image:

  ```bash
  docker build . -t titly-backend
  ```

2. Run the container:

  ```bash
  docker run --rm -p 4000:4000 titly-backend
  ```

## Run the Frontend (SvelteKit)

  1. Navigate to the `titly` directory:

  ```bash
  cd titly
  ```

  2. Install dependencies:

  ```bash
  pnpm install
  ```

  3. Start the development server:

  ```bash
  pnpm run dev
  ```

  4. Open your browser and go to `http://localhost:5173` to view the application.


![image](https://github.com/user-attachments/assets/197a892f-9376-4958-896e-f6e7e5b416b9)

