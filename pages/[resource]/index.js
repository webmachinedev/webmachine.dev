import { useRouter } from "next/router";

export default function Index() {
  const { resource, id } = useRouter().query;
  return (
    <div>
      <p>{resource}</p>
    </div>
  );
}
