document.addEventListener("DOMContentLoaded", function () {
    fetch('/static/questions.json')
        .then(response => response.json())
        .then(questions => {
            questions = shuffleArray(questions);

            const quizContainer = document.getElementById('quiz');
            let currentQuestionIndex = 0;
            let correctAnswers = 0;
            let incorrectAnswers = 0;
            const totalQuestions = questions.length;

            displayQuestion(questions[currentQuestionIndex]);

            function displayQuestion(question) {
                const questionElement = document.createElement('div');
                questionElement.classList.add('question');
                questionElement.innerHTML = `
                    <h3>${question.Question}</h3>
                    <ul>
                        ${shuffleArray(question.Options).map(option => `<li><input type="radio" name="question${question.ID}" value="${option}">${option}</li>`).join('')}
                    </ul>
                    <button onclick="checkAnswer(${question.ID})">Submit</button>
                    <div id="feedback-${question.ID}" class="alert" style="display: none;"></div>
                `;
                quizContainer.innerHTML = '';
                quizContainer.appendChild(questionElement);
            }

            function displayResult() {
                const resultElement = document.createElement('div');
                resultElement.innerHTML = `
                    <h2>Kuis Selesai!</h2>
                    <p>Total Pertanyaan: ${totalQuestions}</p>
                    <p>Jawaban Benar: ${correctAnswers}</p>
                    <p>Jawaban Salah: ${incorrectAnswers}</p>
                    <h3>Pertanyaan dan Jawaban:</h3>
                `;

                questions.forEach(question => {
                    const questionResult = document.createElement('div');
                    questionResult.innerHTML = `
                        <p><strong>${question.Question}</strong></p>
                        <p>Jawaban Benar: ${question.Answer}</p>
                    `;
                    resultElement.appendChild(questionResult);
                });

                quizContainer.innerHTML = '';
                quizContainer.appendChild(resultElement);
            }

            function displayNextQuestion() {
                currentQuestionIndex++;
                if (currentQuestionIndex < totalQuestions) {
                    displayQuestion(questions[currentQuestionIndex]);
                } else {
                    displayResult();
                }
            }

            window.checkAnswer = function (questionID) {
                const selectedOption = document.querySelector(`input[name="question${questionID}"]:checked`);
                if (!selectedOption) {
                    alert('Please select an answer.');
                    return;
                }

                const data = {
                    question_id: questionID,
                    answer: selectedOption.value
                };

                fetch('/check', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data)
                })
                    .then(response => response.json())
                    .then(result => {
                        const feedbackDiv = document.getElementById(`feedback-${questionID}`);
                        if (result.correct) {
                            feedbackDiv.textContent = 'Correct answer!';
                            feedbackDiv.classList.add('correct');
                            correctAnswers++;
                        } else {
                            feedbackDiv.textContent = 'Incorrect answer.';
                            feedbackDiv.classList.add('incorrect');
                            incorrectAnswers++;
                        }
                        feedbackDiv.style.display = 'block';
                        setTimeout(displayNextQuestion, 2000);
                    })
                    .catch(error => console.error('Error checking answer:', error));
            };

            function shuffleArray(array) {
                for (let i = array.length - 1; i > 0; i--) {
                    const j = Math.floor(Math.random() * (i + 1));
                    [array[i], array[j]] = [array[j], array[i]];
                }
                return array;
            }
        })
        .catch(error => console.error('Error fetching questions:', error));
});

function checkAnswer(questionID) {
    const selectedOption = document.querySelector(`input[name="question${questionID}"]:checked`);
    if (!selectedOption) {
        alert('Please select an answer.');
        return;
    }

    const data = {
        question_id: questionID,
        answer: selectedOption.value
    };

    fetch('/check', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(result => {
            const feedbackDiv = document.getElementById(`feedback-${questionID}`);
            if (result.correct) {
                feedbackDiv.textContent = 'Correct answer!';
                feedbackDiv.classList.add('correct');
            } else {
                feedbackDiv.textContent = 'Incorrect answer.';
                feedbackDiv.classList.add('incorrect');
            }
            feedbackDiv.style.display = 'block';
        })
        .catch(error => console.error('Error checking answer:', error));
}
