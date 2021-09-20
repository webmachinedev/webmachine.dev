import { useRouter } from "next/router";

export default function View() {
  const { resource, id } = useRouter().query;
  return (
    <div>
      <p>{resource} {id}</p>
    </div>
  );
}
