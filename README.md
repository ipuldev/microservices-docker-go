<a name="readme-top"></a>

[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#credits">Credits</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Microservices docker is an example how micraservices is actually work in a simple environment, managing traefik and authentication is basic implementation on this project.It's written by Go flexible to add another services using another language. 

Detail Structure Project :
* Auth Service
  * Database:
    * [![Postgresql][Postgresql.com]][Postgresql-url]
  * Routing:
    * Login
    * Signup
    * Authorize
    * Token
* Product Service
  * Database:
    * [![Mongodb][Mongodb.com]][Mongodb-url]
  * Routing:
    * Insert
    * Detail
    * Update
    * Delete
    * List

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Go.com]][Go-url]
* [![Docker][Docker.com]][Docker-url]
* [![Mongodb][Mongodb.com]][Mongodb-url]
* [![Postgresql][Postgresql.com]][Postgresql-url]


<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Prerequisites
* Docker Engine 20.10.18
* Docker Compose version 2.11.2

### Installation
1. ```docker compose up -d```
<br />
Note make sure your docker version is competible
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage
Testing module
- `cd [module]`
- `go test ./mocks/... -run -w`

For more information 

_Postman [Documentation](https://www.postman.com/belivine/workspace/microservices-docker)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Auth Service
- [x] Product Service
- [x] Environment
    - [x] Routing
    - [x] Docker Compose & Image
    - [x] API Gateway
    - [x] Middleware Auth Forward
- [x] Unit Testing
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/[FEATURE]`)
3. Commit your Changes (`git commit -m 'Add some [FEATURE]'`)
4. Push to the Branch (`git push origin feature/[FEATURE]`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Credits
Originally built by [mmorejon](https://github.com/mmorejon/microservices-docker-go-mongodb).

<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[license-shield]: https://img.shields.io/github/license/briankliwon/microservices-docker
[license-url]: https://github.com/briankliwon/microservices-docker/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/muhammad-saiful-abdulah-079545186/
[Design-url]:https://www.figma.com/file/JrLuk8F8CwUI7QUBTkSh3G/Microservices-Docker?node-id=0%3A1
[product-screenshot]: images/screenshot.png
[GO.com]:https://img.shields.io/badge/GO%201.19-0769AD?style=flat&logo=go&logoColor=white
[GO-Url]:https://go.dev/
[Docker.com]:https://img.shields.io/badge/docker-003f8c?style=flat&logo=docker&logoColor=white
[Docker-Url]:https://www.docker.com/
[Mongodb.com]:https://img.shields.io/badge/mongodb-116149?style=flat&logo=mongodb&logoColor=white
[Mongodb-Url]:https://www.mongodb.com/
[Postgresql.com]:https://img.shields.io/badge/postgresql-32658f?style=flat&logo=postgresql&logoColor=white
[Postgresql-Url]:https://www.postgresql.org/
[Redis.com]:https://img.shields.io/badge/redis-d82c20?style=flat&logo=redis&logoColor=white
[Redis-Url]:https://redis.io/
