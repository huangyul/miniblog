import React, { useState } from "react"
import "./index.css"

export interface Todo {
  label: string
  value: string
}

interface ToDoListProps {
  arr: Todo[]
}

const ToDoList: React.FC<ToDoListProps> = ({ arr }) => {

  const [todos, setTodos] = useState<Todo[]>(arr)

  const [inputText, setInputText] = useState<string>('')

  const addTodo = () => {
    if (inputText.trim() != "") {
      const updateTodos = {
        label: inputText,
        value: Date.now().toString()
      }
      setTodos([...todos, updateTodos])
      setInputText('')
    }
  }

  return (
    <>
      <input type="text" value={inputText} onChange={(e) => setInputText(e.target.value)}></input>
      <button type="button" onClick={addTodo}>add</button>
      <div>
        <ul>
          {todos.map(t => {
            return <li key={t.value}>{t.label}</li>
          })}
        </ul>
      </div>
    </>
  )

}

export default ToDoList
