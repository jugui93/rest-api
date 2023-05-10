pipeline {
    agent any

    environment {
        DB_USER = "juan1234"
        DB_PASSWORD = "juan1234"
        DB_NAME = "facts"
        DB_NAME_TEST = "test"
    }

    stages {
        stage('Build') {
            steps {
                // Get some code from a GitHub repository
                git branch: 'main', credentialsId: 'c3901aa1-c7bc-42f7-819e-3cc7219596d7', url: 'git@github.com:jugui93/rest-api.git'

                //Build services
                sh 'docker compose build'
            }
        }
        stage('Start Services') {
            steps {
                // Start all services defined in docker-compose.yml
                script {
                    def composeCommand = "docker compose up -d"
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
                    def composeCommand = "docker compose exec -T web go test ./..."
                    sh composeCommand
                }
            }
        }
    }
}