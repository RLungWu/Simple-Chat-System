# A Simple Chat System
In this system, we use React and Golang to build a distributed chat system.
It's easy and flexible.

## Features
* Real-time Communication: Instantly exchange messages between two separate React applications.
* Built with Go Channels: Utilizes the concurrency mechanisms of Go to ensure efficient and isolated communication.
* Easy Integration: Designed to be easily integrated into existing React applications.
* Scalable Architecture: Engineered for scalability, making it easy to expand beyond two users.


## For the Frontend
```
cd frontend
npm run start
```

## For the Backend
```
docker build -t backend .
docker run -it -p 8080:8080 backend
```

## Demo Pictures for Frontend
![Frontend](https://hackmd.io/_uploads/BkmIy_gkR.png)

## Demo Pictures for Backend
![Backend](https://hackmd.io/_uploads/rk-mx_x1C.png)


### Description
As you can see, the two React App can chat with other. That's because I use the golang's channel to make them more efficient and do not disturb with each other.
