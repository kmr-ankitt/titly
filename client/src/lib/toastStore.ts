import { writable } from 'svelte/store';

export interface ToastMessage {
  id: string;
  type: 'success' | 'error' | 'info';
  message: string;
}

const createToastStore = () => {
  const { subscribe, update } = writable<ToastMessage[]>([]);

  return {
    subscribe,
    add: (message: string, type: 'success' | 'error' | 'info' = 'info', duration = 3000) => {
      const id = Math.random().toString(36).substring(2, 9);
      update((toasts) => [...toasts, { id, type, message }]);

      if (duration > 0) {
        setTimeout(() => {
          update((toasts) => toasts.filter((t) => t.id !== id));
        }, duration);
      }
    },
    remove: (id: string) => {
      update((toasts) => toasts.filter((t) => t.id !== id));
    }
  };
};

export const toast = createToastStore();
