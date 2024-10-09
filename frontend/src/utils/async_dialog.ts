import type {
  DialogOptions,
  DialogReactive,
} from "naive-ui/es/dialog/src/DialogProvider";

type AsyncDialogOptions<T> = DialogOptions & {
  positiveValue?: T;
  negativeValue?: T;
  maskValue?: T;
};

export function asyncDialog<T>(
  d: (options: DialogOptions) => DialogReactive,
  option: AsyncDialogOptions<T>
): Promise<T | null> {
  return new Promise<T | null>((resolve) => {
    d({
      ...option,
      onPositiveClick() {
        resolve(option.positiveValue ?? null);
      },
      onNegativeClick() {
        resolve(option.negativeValue ?? null);
      },
      onMaskClick() {
        resolve(option.maskValue ?? null);
      },
    });
  });
}
