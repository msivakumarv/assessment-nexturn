document.addEventListener('DOMContentLoaded', () => {
  const newTaskInput = document.getElementById('new-task');
  const addTaskButton = document.getElementById('add-task');
  const taskList = document.getElementById('task-list');
  const pendingCountElement = document.getElementById('pending-count');


  let tasks = JSON.parse(localStorage.getItem('tasks')) || [];

  function updateLocalStorage() {
      localStorage.setItem('tasks', JSON.stringify(tasks));
  }

  function updatePendingCount() {
      const pendingCount = tasks.filter(task => !task.completed).length;
      pendingCountElement.textContent = `Pending Tasks: ${pendingCount}`;
  }

  function renderTasks() {
      taskList.innerHTML = '';
      tasks.forEach((task, index) => {
          const li = document.createElement('li');
          li.className = `task ${task.completed ? 'completed' : ''}`;
          
          li.innerHTML = `
              <span class="task-name">${task.name}</span>
              <div>
                  <button class="edit">Edit</button>
                  <button class="delete">Delete</button>
              </div>
          `;

          li.addEventListener('click', () => toggleComplete(index));

          
          li.querySelector('.edit').addEventListener('click', (event) => {
              event.stopPropagation();
              editTask(index);
          });

          li.querySelector('.delete').addEventListener('click', (event) => {
              event.stopPropagation(); 
              deleteTask(index);
          });

          taskList.appendChild(li);
      });
      updatePendingCount();
  }

  function addTask(taskName) {
      if (taskName.trim() === '') {
          alert('Task cannot be empty!');
          return;
      }
      tasks.push({ name: taskName, completed: false });
      updateLocalStorage();
      renderTasks();
      newTaskInput.value = '';
  }

  function editTask(index) {
      const newTaskName = prompt('Edit Task:', tasks[index].name);
      if (newTaskName !== null && newTaskName.trim() !== '') {
          tasks[index].name = newTaskName.trim();
          updateLocalStorage();
          renderTasks();
      }
  }

  function deleteTask(index) {
      tasks.splice(index, 1);
      updateLocalStorage();
      renderTasks();
  }

  function toggleComplete(index) {
      tasks[index].completed = !tasks[index].completed;
      updateLocalStorage();
      renderTasks();
  }

  
  addTaskButton.addEventListener('click', () => addTask(newTaskInput.value));
  newTaskInput.addEventListener('keypress', (e) => {
      if (e.key === 'Enter') {
          addTask(newTaskInput.value);
      }
  });


  renderTasks();
});
