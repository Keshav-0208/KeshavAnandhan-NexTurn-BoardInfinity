function addTask() {
    var taskInput = document.getElementById("taskInput");
    var taskValue = taskInput.value.trim();

    if (taskValue) {
        var taskList = document.getElementById("taskList");
        var newTask = document.createElement("li");

        newTask.innerHTML = `
            <input type="checkbox" onclick="finishTask(this)">
            <span>${taskValue}</span>
            <span onclick="removeTask(this)" style="cursor: pointer; color: red; margin-left: 10px;">x</span>
        `;

        taskList.appendChild(newTask);
        taskInput.value = "";
    }
}

function removeTask(task) {
    var taskList = document.getElementById("taskList");
    taskList.removeChild(task.parentElement);
}

function finishTask(checkbox) {
    var taskText = checkbox.nextElementSibling;
    if (checkbox.checked) {
        taskText.style.textDecoration = "line-through";
    } else {
        taskText.style.textDecoration = "none";
    }
}
