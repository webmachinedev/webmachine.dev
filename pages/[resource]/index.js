import { useRouter } from "next/router";
import { useEffect, useState } from "react";

export default function Index() {
  const { resource } = useRouter().query;

  const [error, setError] = useState()
  const [items, setItems] = useState([]);

  useEffect(() => {
    (async () => {
        const {res, err} = await callFunction(`github.com/webmachinedev/go-clients/db/${resource}`, "Index")
        setError(err)
        setItems(res)
    })();
  }, [resource]);

  if (error) return error;

  return (
    <div>
      <p>{resource}</p>
      <br />
      <ul>
        {items.map((item) => (
          <li key={item.id}>
            <a href={`/${resource}/${item.id}`}>{item.name || item.id}</a>
          </li>
        ))}
      </ul>
    </div>
  );
}


async function callFunction(pkg, func, args = []) {
    const res = await fetch("https://webmachine.dev/api", {
        method: "POST",
        body: JSON.stringify({pkg, func, args})
    })
    return await res.json();
}
