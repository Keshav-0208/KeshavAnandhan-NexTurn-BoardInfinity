// script.js
const quizData = [
    {
        question: "What is the capital of France?",
        options: ["Berlin", "Madrid", "Paris", "Rome"],
        correct: "Paris"
    },
    {
        question: "Which planet is known as the Red Planet?",
        options: ["Earth", "Mars", "Jupiter", "Saturn"],
        correct: "Mars"
    },
    {
        question: "Who wrote 'Hamlet'?",
        options: ["Shakespeare", "Dickens", "Tolstoy", "Hemingway"],
        correct: "Shakespeare"
    }
];

function loadQuiz() {
    const quizContainer = document.getElementById('quiz');
    quizData.forEach((item, index) => {
        const questionElement = document.createElement('div');
        questionElement.innerHTML = `<h3>${item.question}</h3>`;
        
        item.options.forEach(option => {
            const optionElement = document.createElement('div');
            optionElement.innerHTML = `
                <input type="radio" name="question${index}" value="${option}">
                <label>${option}</label>
            `;
            questionElement.appendChild(optionElement);
        });
        
        quizContainer.appendChild(questionElement);
    });
}

function submitQuiz() {
    let score = 0;
    
    quizData.forEach((item, index) => {
        const selectedOption = document.querySelector(`input[name="question${index}"]:checked`);
        
        if (selectedOption && selectedOption.value === item.correct) {
            score++;
        }
    });
    
    document.getElementById('score').innerText = `Your score: ${score}/${quizData.length}`;
}

loadQuiz();
