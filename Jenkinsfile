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
        stage('Test') {
            steps {
                //Create and start containers
                sh 'docker compose up'

                //Run test
                sh 'docker compose exec web go test ./...'
                //Stop containers
                sh 'docker compose stop'
            }
        }
    }
}