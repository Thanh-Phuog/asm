pipeline {
    agent any

    environment {
        DOCKER_IMAGE_SERVER = 'thahphuog/devopsasm'
        DOCKER_IMAGE_CLIENT = 'thahphuog/client'
        DOCKER_TAG = '1.0.0'
        TELEGRAM_BOT_TOKEN = '7908085505:AAEy0dz1yrVesOaFmZ1s5qWlvslKWekBi_k'
        TELEGRAM_CHAT_ID = '-1002403309943'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: 'https://github.com/Thanh-Phuog/asm.git'
            }
        }

        stage('Build Client Docker Image') {
            steps {
                script {
                    echo 'Building Client Docker image'
                    docker.build("${DOCKER_IMAGE_CLIENT}:${DOCKER_TAG}", './book_mana_client')
                }
            }
        }

        stage('Build Server Docker Image') {
            steps {
                script {
                    echo 'Building Server Docker image'
                    docker.build("${DOCKER_IMAGE_SERVER}:${DOCKER_TAG}", './book_mana_server')
                }
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    // Push Server image
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE_SERVER}:${DOCKER_TAG}").push()
                    }

                    // Push Client image
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE_CLIENT}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                script {
                    echo 'Clearing server_golang-related images and containers...'
                    sh '''
                        docker container stop server-golang || echo "No container named server-golang to stop"
                        docker container rm server-golang || echo "No container named server-golang to remove"
                        docker image rm ${DOCKER_IMAGE_SERVER}:${DOCKER_TAG} || echo "No image ${DOCKER_IMAGE_SERVER}:${DOCKER_TAG} to remove"
                    '''

                    echo 'Deploying to DEV environment...'
                    sh '''
                        docker image pull ${DOCKER_IMAGE_SERVER}:${DOCKER_TAG}
                        docker network create dev || echo "Network already exists"
                        docker container run -d --rm --name server-golang -p 3000:3000 --network dev ${DOCKER_IMAGE_SERVER}:${DOCKER_TAG}
                    '''
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }

        success {
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}

def sendTelegramMessage(String message) {
    sh """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
    -d chat_id=${TELEGRAM_CHAT_ID} \
    -d text="${message}"
    """
}
