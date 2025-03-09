# Cosmic Gate Online Web Chat

## Description:
Cosmic Gate Online Web Chat is a real-time messaging web application allowing users to create profiles, find other users, add friends, and engage in seamless, live conversations.
The app leverages WebSockets for real-time communication and MongoDB for storing user and message data.
Designed for simplicity and ease of use, this chat application ensures all users a smooth and interactive experience.

## Features:
Create a User Profile: Set up a personalized profile with a username, bio, and more.
Find Users: Search for users within the platform to connect with.
Add Friends: Send and receive friend requests to build your network.
Real-Time Chat: Send and receive messages instantly with WebSocket technology.
User Authentication: Secure login and registration using email and password.

## Project Structure:

### Frontend (Client)
**Path:** ./client

**Technology:** Next.js (React framework)

**Real-Time Messaging:** WebSocket

**How to Run:**
1. Navigate to the ./client directory.
2. Install dependencies:
```
npm install
```
3. Start the development server:
```
npm run dev
```
4. Open your browser and go to http://localhost:3000 to access the app.

### Backend (Server)
**Path:** ./server

**Technology:** Go (Golang), WebSocket, MongoDB

**How to Run:**
1. Navigate to the ./server directory.
2. Install MongoDB and set up the database.
3. Run the server:
```
go run .
```
4. The server will start, and the frontend will be able to interact with it for real-time chat.
