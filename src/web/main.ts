import focus from "@alpinejs/focus";
import Alpine from "alpinejs";

Alpine.data("data", (id: string) => {
  const raw = document.querySelector(`script[type="application/json"]#${id}`)
    ?.textContent ?? "{}";
  return JSON.parse(raw);
});

Alpine.plugin(focus);

Alpine.start();
