import ToDoList, { Todo } from "@/page/todo-list"

function App() {


  const todos: Todo[] = [
    {
      label: "1",
      value: "1"
    }
  ]

  return (
    <div>
      <ToDoList arr={todos}></ToDoList>
    </div>
  )
}

export default App
