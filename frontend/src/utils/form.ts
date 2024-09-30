import type { FormInst } from "naive-ui";

export async function validateForm(
  form: FormInst | null | undefined
): Promise<boolean> {
  if (!form) return false;
  try {
    await form.validate();
    return true;
  } catch (error) {
    return false;
  }
}
