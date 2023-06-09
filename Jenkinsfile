pipeline {
  agent {
    label 'ubuntu'
  }
  environment {
    DB_USER = "${DB_USER}"
    DB_PASSWORD = "${DB_PASSWORD}"
    DB_NAME = "${DB_NAME}"
    DB_NAME_TEST = "${DB_NAME_TEST}"
  }
  stages {
    stage('SCM'){
      steps{
        git(branch: 'main', credentialsId: 'c3901aa1-c7bc-42f7-819e-3cc7219596d7', url: 'git@github.com:jugui93/rest-api.git')
      }
    }
    stage('SonarQube analysis'){
      steps{
        script{
          def scannerHome = tool 'sq1';
          withSonarQubeEnv('SonarQube') { // If you have configured more than one global server connection, you can specify its name
          sh "${scannerHome}/bin/sonar-scanner"
          }
        }
      }
    }
    stage("Quality Gate") {
      steps {
        retry(3){
          timeout(time: 15, unit: 'SECONDS') {
              // Parameter indicates whether to set pipeline to UNSTABLE if Quality Gate fails
              // true = set pipeline to UNSTABLE, false = don't
              waitForQualityGate abortPipeline: true
          }
        }
      }
    }
    stage('Build Test') {
      steps {
        git(branch: 'main', credentialsId: 'c3901aa1-c7bc-42f7-819e-3cc7219596d7', url: 'git@github.com:jugui93/rest-api.git')
        sh 'docker compose -f docker-compose.test.yml build'
      }
    }

    stage('Start Services') {
      steps {
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
      steps {
        sh '''aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 181021887246.dkr.ecr.us-east-1.amazonaws.com
                docker compose build
                docker tag project-lab-app-web:latest 181021887246.dkr.ecr.us-east-1.amazonaws.com/project-lab:latest
                docker push 181021887246.dkr.ecr.us-east-1.amazonaws.com/project-lab:latest'''
      }
    }

    stage('Clean Docker') {
      steps {
        sh 'docker system prune --all --force'
      }
    }

    stage('Deploy') {
      agent any
      steps {
        sh 'ssh -o StrictHostKeyChecking=no ubuntu@54.81.202.196 "aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 181021887246.dkr.ecr.us-east-1.amazonaws.com"'
        sh 'ssh -o StrictHostKeyChecking=no ubuntu@54.81.202.196 "docker pull 181021887246.dkr.ecr.us-east-1.amazonaws.com/project-lab:latest"'
        sh 'ssh -o StrictHostKeyChecking=no ubuntu@54.81.202.196 "docker tag 181021887246.dkr.ecr.us-east-1.amazonaws.com/project-lab project-lab-app-web"'
        script {
          sh  'ssh -o StrictHostKeyChecking=no ubuntu@54.81.202.196 "cd /home/ubuntu/project-lab && docker compose up -d"'

          // Wait for the web service to be ready
          retry(5) {
            def response = sh(returnStdout: true, script: 'ssh -o StrictHostKeyChecking=no ubuntu@54.81.202.196 "curl -s -o /dev/null -w "%{http_code}" http://localhost:3000/fact"')
            if (response.trim() == '200') {
              echo 'Web service is ready!'
            } else {
              echo 'Web service is not ready yet, retrying after a delay...'
              sleep time: 8, unit: 'SECONDS'  // Add a delay of 30 seconds
              error 'Web service is not ready yet'
            }
          }
        }

      }
    }

  }
}