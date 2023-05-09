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

                // Run Maven on a Unix agent.
                step([$class: 'DockerComposeBuilder', dockerComposeFile: 'docker-compose.yml', option: [$class: 'StartAllServices'], useCustomDockerComposeFile: true])

                // To run Maven on a Windows agent, use
                // bat "mvn -Dmaven.test.failure.ignore=true clean package"
            }

        }
    }
}