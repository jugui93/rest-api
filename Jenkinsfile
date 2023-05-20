pipeline {
    agent {
        label 'ubuntu'
    }

    environment {
        DB_USER = "juan1234"
        DB_PASSWORD = "juan1234"
        DB_NAME = "facts"
        DB_NAME_TEST = "test"
    }

    stages {
        stage('Build Test') {
            steps {
                // Get some code from a GitHub repository
                git branch: 'main', credentialsId: 'c3901aa1-c7bc-42f7-819e-3cc7219596d7', url: 'git@github.com:jugui93/rest-api.git'

                //Build services
                sh 'docker compose -f docker-compose.test.yml build'
            }
        }
        stage('Start Services') {
            steps {
                // Start all services defined in docker-compose.yml
                script {
                    def composeCommand = "docker compose -f docker-compose.test.yml up -d"
                    sh composeCommand

                    // Wait for the web service to be ready
                    retry(5) {
                        def response = sh(returnStdout: true, script: 'curl -s -o /dev/null -w "%{http_code}" http://localhost:3000/fact')
                        if (response.trim() == '200') {
                            echo 'Web service is ready!'
                        } else {
                            error 'Web service is not ready yet'
                        }
                    }
                }
            }
        }
        stage('Run Test') {
            steps {
                // Run tests inside the web service container
                script {
                    def testResult = sh(script: 'docker compose exec -T web go test ./cmd', returnStatus: true)
                    if (testResult != 0) {
                        currentBuild.result = 'FAILURE'
                        sh 'docker compose -f docker-compose.test.yml down -v'
                        error('Tests failed')
                    }
                    sh 'docker compose -f docker-compose.test.yml down -v'
                }
            }
        }
        stage('Build and push Docker compose') {
            steps{
                // Build your Docker Compose here
                sh '''aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 181021887246.dkr.ecr.us-east-1.amazonaws.com
                docker compose build
                docker tag project-lab-app-web:latest 181021887246.dkr.ecr.us-east-1.amazonaws.com/project-lab:latest
                docker push 181021887246.dkr.ecr.us-east-1.amazonaws.com/project-lab:latest'''
            }
        }
        stage('Clean Docker') {
            steps {
                // Clean Docker
                sh 'docker system prune --all --force'
            }
        }
        stage('Deploy') {
            steps {
                sshagent(['61b93c39-4d3b-4f7d-a4a4-cbe4d7597894']){
                    // Pull the latest image from ECR
                    sh 'ssh ubuntu@54.81.202.196 "aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 181021887246.dkr.ecr.us-east-1.amazonaws.com"'
                    sh 'ssh ubuntu@54.81.202.196 "docker pull 181021887246.dkr.ecr.us-east-1.amazonaws.com/repository:latest"'

                    // Deploy the app using Docker Compose
                    script {
                        sh  'ssh ubuntu@54.81.202.196 "docker compose up -d"'

                        // Wait for the web service to be ready
                        retry(5) {
                            def response = sh(returnStdout: true, script: 'curl -s -o /dev/null -w "%{http_code}" http://localhost:3000/fact')
                            if (response.trim() == '200') {
                                echo 'Web service is ready!'
                            } else {
                                error 'Web service is not ready yet'
                            }
                        }
                    }
                }
            }
        }
    }
}