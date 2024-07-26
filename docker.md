# Docker Docs
-------------
## Content Course
### Section 1: 
    Problem, What and Why Docker?
### Section 2:
    Dockerfile
	Images and Containers
	Docker Commands
	
### Section 3:
	Local Development and Hot-Reload
	Volumes
	Docker Compose
	Environments (Dev, Staging, and Production)
	
### Section 4:
	Docker with DB

### Section 5:
	Production Deployment
	Load balancing with Nginx
	Automation with Watchtower
	How CI/CD works?

### Section 6:
	Docker Orchestration (Kubernetes, Swarm, ...etc)
	Docker Swarm
	Rolling Updates
	
	
## Section 1: Problem, What and Why Docker?
-------------------------------------------
	Problem!
	--------
	A new developer is joining a company
		Backend, frontend
		installation Errors
		version Downgrade!... 
	Deploy your app
		Testing, Staging, and Production
		
## Section 2: Dockerfile Images and Containers, Docker Commands
---------------------------------------------------------------
	
### Dockerfile
	
	Dockerfile
    ----------
		//FROM baseImage
		FROM node:14
		WORKDIR /app
		//COPY package.json /app
		COPY package.json .
		//RUN command
		RUN npm install
		//COPY index.js .
		COPY . .
		
		EXPOSE 4000
		
		//CMD ["executable"]
		CMD ["npm", "start"]
		
	Docker Commands
	---------------
		docker build ./Dockerfile
		docker build .  //>> . : Dockerfile
		docker build -t express-node-app . // . : Dockerfile, -t express-node-app : name the build image
			==>> //Build Image
			==>> Dockerfile >> Build Image >> Run Container
		
		//Show images
        -------------
		docker image ls
			> express-node-app
			
		//Run Container
		docker run express-node-app
		docker run --name express-node-app-container express-node-app
		
		//Show docker
        -------------
		docker ps
			CONTAINER ID	IMAGE		COMMAND   CREATED    STATUS    PORTS     NAME
			083abb45fd78    express-node-app			       4000/tcp	 express-node-app-container
										0.0.0.0:4000->4000/tcp
		//Delete docker
        ---------------
		docker rm express-node-app-container -f
		
		//Stop docker
        -------------
		docker stop express-node-app-container
		
		//Close Terminal with Run container : -d  detach
        ----------------------------------------
		docker run --name express-node-app-container -d express-node-app
		
		//Use Port
        ----------
		docker run --name express-node-app-container -d -p 4000:4000 express-node-app
		
		
		//How to Run Docker
        -------------------
		==>> Dockerfile >> Build Image >> Run Container 1, Run Container 2
		
		//Docker pull
        -------------
		docker pull php
		
		//Execute Terminal "Shell" in Container
        ---------------------------------------
		docker exec express-node-app-container bash
		docker exec -it express-node-app-container bash
			>> root@kjdlfjdfds:/app# pwd
			   /app
			>> exit
			   
		//.dockerignore
        ---------------
		.dockerignore
            /node_modules
			Dockerfile
			.env
			docker-compose*
		
		
## Section 3: Local Development and Hot-Reload
-------------------------------------------
	FROM node:14
	
	WORKDIR /app
	
	COPY package.json .
	
	RUN npm install
	
	COPY . .
	
	EXPOSE 4000
	
	CMD ["npm", "run", "start-dev"] // ==> Hot-Reload
	
	
    //Hot-Reload - Volumes  // Volumes <==> DataBase
    ----------------------
	docker rm express-node-app-container -f
	docker build -t express-node-app .
    docker run --name express-node-app-container -d -p 4000:4000 express-node-app

	docker run --name express-node-app-container "copy path the project folder" -v /Users/...../my-express-app:/app:ro -d -p 4000:4000 express-node-app

    docker run --name express-node-app-container "copy path the project folder" -v $(pwd):/app:ro -d -p 4000:4000 express-node-app

    docker run --name express-node-app-container -v /app/node_modules -d -p 4000:4000 express-node-app

    docker run --name express-node-app-container -v $(pwd):/app -v /app/node_modules -d -p 4000:4000 express-node-app

    docker run --name express-node-app-container -v $(pwd)/src:/app/src:ro -d -p 4000:4000 express-node-app

	
	//Docker Logs
    -------------
	docker logs express-node-app-container
	
	//Shell Docker
    --------------
	docker exec -it express-node-app-container bash

    //Used Types
    ------------
        Bind mount
        Anonymous


    Docker Compose
    --------------
        docker --version ## docker-compose --version
        docker-compose --help
        docker --help
        docker run --help
            >> -e, --env list

        docker-compose up -d
        docker-compose down

        //Show Containers
        docker ps

        docker run --name express-node-app-container -v $(pwd)/src:/app/src:ro -d -p 4000:4000 express-node-app

        //Create File : docker-compose.yml
        version: "3"
        services:
            node-app:   //name the app
                container_name: express-node-app-container
                build: .
                volumes:
                    - ./src:/app/src:ro
                ports:
                    - "4000:4000"
                environment:  //can remplace with env_file:
                    - PORT=4000
                    - NODE_ENV=production   
                env_file:
                    - ./.env  // filename


    Environment Variables
    ---------------------
        >> index.js
            const PORT = process.env.PORT || 4000;
        >> Dockerfile
            ENV PORT=4000
            EXPOSE $PORT

        docker run --name express-node-app-container -v $(pwd)/src:/app/src:ro --env PORT=4000 --env NODE_ENV=development -d -p 4000:4000 express-node-app

        docker exec -it express-node-app-container bash
            >> printenv

        //Create File Environment : .env
        .env
            PORT=4000
            NODE_ENV=development
            DB_HOST=12345
            DB_PASSWORDO=password

        docker run --name express-node-app-container -v $(pwd)/src:/app/src:ro --env-file ./.env -d -p 4000:4000 express-node-app

        //Create Files the Environment
        docker-compose.dev.yml
        docker-compose.ref.yml
        docker-compose.prod.yml

        docker-compose.prod.yml
            version: "3"
            services:
                node-app:
                    container_name: express-node-app-container
                    build: .
                    ports:
                        - "4000:4000"   
                    environment:
                        - NODE_ENV=production
                    env_file:
                        - ./.env  // filename

        docker-compose.dev.yml
            version: "3"
            services:
                node-app:   //name the app
                    container_name: express-node-app-container
                    build: .
                    volumes:
                        - ./src:/app/src:ro
                    ports:
                        - "4000:4000"
                    environment:
                        - NODE_ENV=development
                    env_file:
                        - ./.env  // filename

        
        .env
            PORT=4000
            DB_HOST=12345
            DB_PASSWORD=password

        docker-compose up -d
            ERROR
        
        docker-compose -f docker-compose.dev.yml up -d
        docker-compose -f docker-compose.dev.yml down

        docker-compose -f docker-compose.prod.yml up -d

        Create docker-compose and docker-compse dev and prod
        ----------------------------------------------------

         docker-compose.yml
         ------------------
            version: "3"
            services:
                node-app:
                    container_name: express-node-app-container
                    build: .
                    ports:
                        - "4000:4000"
                    env_file:
                        - ./.env 

         docker-compose.dev.yml
         ----------------------
            version: "3"
            services:
                node-app:  
                    volumes:
                        - ./src:/app/src:ro
                    environment:
                        - NODE_ENV=development
      
         docker-compose.prod.yml
         ----------------------
            version: "3"
            services:
                node-app:  
                    volumes:
                        - ./src:/app/src:ro
                    environment:
                        - NODE_ENV=production

        docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
        docker-compose -f docker-compose.yml -f docker-compose.dev.yml down

        docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
        docker-compose -f docker-compose.yml -f docker-compose.prod.yml down

        docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build