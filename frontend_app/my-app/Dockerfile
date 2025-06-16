# Use the latest LTS version of Node.js
FROM node:24-alpine
 
# Set the working directory inside the container
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
 
# Expose the port your app runs on
EXPOSE 3000
 
# TODO: Define the command to run your app. 
CMD ["npm", "run", "dev"]