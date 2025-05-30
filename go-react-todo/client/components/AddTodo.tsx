import { useState } from "react";
import { useForm } from "@mantine/form";
import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";
import { ENDPOINT, type Todo } from "../src/App";
import type { KeyedMutator } from "swr";

function AddTodo({ mutate }: { mutate: KeyedMutator<Todo[]>}) {
  const [open, setOpen] = useState(false);

  const form = useForm({
    initialValues: {
      title: "",
      body: "",
    },
  });

//   async function createTodo(values: { title: string; body: string }) {
//     const updated = await fetch(`${ENDPOINT}/api/todos`, {
//       method: "POST",
//       headers: {
//         "Content-Type": "application/json",
//       },
//       body: JSON.stringify(values),
//     }).then((r) => r.json());

//       mutate(updated);
//     form.reset();
//     setOpen(false);
    //   }
    
    async function createTodo(values: { title: string; body: string }) {
      try {
        const updated = await fetch(`${ENDPOINT}/api/todos`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(values),
        }).then((r) => {
          if (!r.ok) throw new Error(`Status ${r.status}`);
          return r.json();
        });

        mutate(updated);
        form.reset();
        setOpen(false);
      } catch (err) {
        alert("Error creating todo: " + err);
        console.error(err);
      }
    }
      

  return (
    <>
      <Modal opened={open} onClose={() => setOpen(false)} title="Create todo">
        <form onSubmit={form.onSubmit(createTodo)}>
          <TextInput
            required
            mb={12}
            label="Todo"
            placeholder="What do you want to do?"
            {...form.getInputProps("title")}
          />
          <Textarea
            required
            mb={12}
            label="Body"
            placeholder="Tell me more..."
            {...form.getInputProps("body")}
          />
          <Button type="submit">Create Todo</Button>
        </form>
      </Modal>

      <Group justify="center">
        <Button
          fullWidth
          mb={12}
          onClick={() => setOpen(true)}
          className="add-todo-button"
        >
          ADD TODO
        </Button>
      </Group>
    </>
  );
}

export default AddTodo;
