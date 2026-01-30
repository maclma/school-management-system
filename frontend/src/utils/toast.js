/**
 * Toast/Notification System
 * Provides centralized notification management with auto-dismiss
 */

export class Toast {
  static instance = null;
  static listeners = [];

  constructor() {
    this.notifications = [];
    this.nextId = 1;
  }

  static getInstance() {
    if (!Toast.instance) {
      Toast.instance = new Toast();
    }
    return Toast.instance;
  }

  static subscribe(listener) {
    Toast.listeners.push(listener);
    return () => {
      Toast.listeners = Toast.listeners.filter(l => l !== listener);
    };
  }

  static notify(listeners) {
    Toast.listeners.forEach(listener => listener(Toast.getInstance().notifications));
  }

  /**
   * Show success notification
   */
  static success(message, duration = 3000) {
    return Toast.getInstance().add({
      type: 'success',
      message,
      duration,
    });
  }

  /**
   * Show error notification
   */
  static error(message, duration = 5000) {
    return Toast.getInstance().add({
      type: 'error',
      message,
      duration,
    });
  }

  /**
   * Show warning notification
   */
  static warning(message, duration = 4000) {
    return Toast.getInstance().add({
      type: 'warning',
      message,
      duration,
    });
  }

  /**
   * Show info notification
   */
  static info(message, duration = 3000) {
    return Toast.getInstance().add({
      type: 'info',
      message,
      duration,
    });
  }

  /**
   * Show loading notification (no auto-dismiss)
   */
  static loading(message) {
    return Toast.getInstance().add({
      type: 'loading',
      message,
      duration: 0,
    });
  }

  add(notification) {
    const id = this.nextId++;
    const toast = {
      id,
      ...notification,
      dismiss: () => this.remove(id),
    };

    this.notifications.push(toast);
    Toast.notify();

    if (notification.duration > 0) {
      setTimeout(() => this.remove(id), notification.duration);
    }

    return toast;
  }

  remove(id) {
    this.notifications = this.notifications.filter(n => n.id !== id);
    Toast.notify();
  }

  clear() {
    this.notifications = [];
    Toast.notify();
  }

  getAll() {
    return [...this.notifications];
  }
}

/**
 * React Hook for Toast notifications
 */
export function useToast() {
  const [notifications, setNotifications] = React.useState([]);

  React.useEffect(() => {
    const unsubscribe = Toast.subscribe(setNotifications);
    return unsubscribe;
  }, []);

  return {
    success: Toast.success,
    error: Toast.error,
    warning: Toast.warning,
    info: Toast.info,
    loading: Toast.loading,
    dismiss: (id) => Toast.getInstance().remove(id),
    clear: () => Toast.getInstance().clear(),
    notifications,
  };
}

/**
 * Promise wrapper with loading toast
 */
export async function toastAsync(promise, messages = {}) {
  const {
    loading: loadingMsg = 'Loading...',
    success: successMsg = 'Success!',
    error: errorMsg = 'Something went wrong',
  } = messages;

  const loadingToast = Toast.loading(loadingMsg);

  try {
    const result = await promise;
    loadingToast.dismiss();
    Toast.success(successMsg);
    return result;
  } catch (err) {
    loadingToast.dismiss();
    Toast.error(errorMsg);
    throw err;
  }
}

/**
 * Toast notification types with colors
 */
export const toastConfig = {
  success: {
    bg: '#10b981',
    icon: '✓',
    borderColor: '#059669',
  },
  error: {
    bg: '#ef4444',
    icon: '✕',
    borderColor: '#dc2626',
  },
  warning: {
    bg: '#f59e0b',
    icon: '⚠',
    borderColor: '#d97706',
  },
  info: {
    bg: '#3b82f6',
    icon: 'ℹ',
    borderColor: '#2563eb',
  },
  loading: {
    bg: '#6366f1',
    icon: '⟳',
    borderColor: '#4f46e5',
  },
};
