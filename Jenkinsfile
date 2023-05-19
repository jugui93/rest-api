pipeline {
    agent {
        label 'ubuntu'
    }

    environment {
        DB_USER = ${DB_USER}
        DB_PASSWORD = ${DB_PASSWROD}
        DB_NAME = ${DB_NAME}
        DB_NAME_TEST = ${DB_NAME_TEST}
        AWS_ACCESS_KEY_ID = credentials('aws-access-key-id')
        AWS_SECRET_ACCESS_KEY = credentials('aws-secret-access-key')
        AWS_DEFAULT_REGION = 'us-east-1'
        ECR_REPOSITORY_URL = '181021887246.dkr.ecr.us-east-1.amazonaws.com/project-lab'
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
        stage('Build') {
            steps{
                // Build your Docker Compose here
                withAmazonECR(credentialsId: '70ba9347-f845-4a24-84ae-e9abb7b28bff', region: 'us-east-1') {
                    def repository = "project-lab"
                    def tag = "latest"
                    ecr.createRepository(repository)
                    sh "docker compose build"
                    sh "docker tag project-lab-app-web:latest ${ecr.registry(repository)}:${tag}"
                    ecr.push(repository: repository, tag: tag)
                }
            }
        }
        // stage('Push') {
        //     steps {
        //         withCredentials([string(credentialsId: 'aws-access-key-id', variable: 'AWS_ACCESS_KEY_ID'), 
        //                         string(credentialsId: 'aws-secret-access-key', variable: 'AWS_SECRET_ACCESS_KEY')]) {
        //         sh "aws ecr get-login-password | docker login --username AWS --password-stdin $ECR_REPOSITORY_URL"
        //         sh "docker push $ECR_REPOSITORY_URL:latest"
        //         }
        //     }
        // }
    }
}