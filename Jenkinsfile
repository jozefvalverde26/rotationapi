pipeline {
    agent any
    tools {
        go ''
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Get Code') {
            steps {
                script {
                  scmVars = checkout scm
                  echo 'scm : the commit id is ' + scmVars.GIT_COMMIT
                }
            }
        }
        stage('Build') {
            steps {
                sh 'go build ./cmd/rotationapi/'
            }
        }
        stage('Run Tests') {
            steps {
                sh 'go test ./... -coverprofile=coverage.txt'
            }
        }
    }
}
