# Product Service

## Overview

This project sets up a basic microservice for managing products using Fiber in Go.

## Project Structure

The folder structure includes directories for configuration, database access, models, handlers, repositories, services, routes, and more.

## Setup Instructions

1. **Clone the Repository**

   \`\`\`bash
   git clone https://github.com/your-org/product-service.git
   cd product-service
   \`\`\`

2. **Create Project Structure**

   Run the setup script to create the directory structure and files:

   \`\`\`bash
   ./setup.sh
   \`\`\`

3. **Install Dependencies**

   \`\`\`bash
   go mod tidy
   \`\`\`

4. **Configure Environment**

   Create a `.env` file in the root directory with:

   \`\`\`
   MONGO_URI=mongodb://username:password@localhost:27017/dbname
   \`\`\`

6. **Dev Run**

   Run the application using air:

   \`\`\`bash
  air -c \ .air.toml
   \`\`\`

5. **Build and Run**

   Build the application:

   \`\`\`bash
   go build -o main ./cmd
   \`\`\`

   Run the application:

   \`\`\`bash
   ./main
   \`\`\`

## License

This project is licensed under the MIT License. See the LICENSE file for details.