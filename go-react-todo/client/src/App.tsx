import './App.css'
import { Box, List } from "@mantine/core";
import useSWR from "swr";
import AddTodo from '../components/AddTodo.tsx'; 

export interface Todo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

export const ENDPOINT = "http://localhost:4000";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {

  const {data, mutate} = useSWR<Todo[]>('api/todos', fetcher)

  return (
    <Box
      sx={(theme) => ({
        padding: "2rem",
        width: "100%",
        maxWidth: "40rem", // ✅ correct casing
        margin: "0 auto",
      })}
    >
      <List spacing="xs" size="sm" mb={12} center>
        {data?.map((todo) => {
          return (
            <List.Item key={`todo_list__${todo.id}`}>{todo.title}</List.Item>
          );
        })}
      </List>

      <AddTodo mutate={mutate} />
    </Box>
  );
}

export default App
