/**
 * Form Validation Utilities
 * Provides comprehensive validation functions for common form fields
 */

export const validators = {
  // Email validation
  email: (email) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!email) return { valid: false, error: 'Email is required' };
    if (!emailRegex.test(email)) return { valid: false, error: 'Invalid email format' };
    return { valid: true };
  },

  // Password validation (minimum 8 chars, uppercase, lowercase, number)
  password: (password) => {
    if (!password) return { valid: false, error: 'Password is required' };
    if (password.length < 8) return { valid: false, error: 'Password must be at least 8 characters' };
    if (!/[A-Z]/.test(password)) return { valid: false, error: 'Password must contain uppercase letter' };
    if (!/[a-z]/.test(password)) return { valid: false, error: 'Password must contain lowercase letter' };
    if (!/[0-9]/.test(password)) return { valid: false, error: 'Password must contain number' };
    return { valid: true };
  },

  // Strong password (current + special character)
  strongPassword: (password) => {
    const basic = validators.password(password);
    if (!basic.valid) return basic;
    if (!/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password)) {
      return { valid: false, error: 'Password must contain special character' };
    }
    return { valid: true };
  },

  // Name validation (2-50 characters, letters and spaces only)
  name: (name) => {
    if (!name) return { valid: false, error: 'Name is required' };
    if (name.trim().length < 2) return { valid: false, error: 'Name must be at least 2 characters' };
    if (name.trim().length > 50) return { valid: false, error: 'Name must not exceed 50 characters' };
    if (!/^[a-zA-Z\s'-]+$/.test(name)) return { valid: false, error: 'Name can only contain letters, spaces, and apostrophes' };
    return { valid: true };
  },

  // Phone number validation (10-15 digits)
  phone: (phone) => {
    if (!phone) return { valid: false, error: 'Phone number is required' };
    const digitsOnly = phone.replace(/\D/g, '');
    if (digitsOnly.length < 10) return { valid: false, error: 'Phone number must have at least 10 digits' };
    if (digitsOnly.length > 15) return { valid: false, error: 'Phone number must not exceed 15 digits' };
    return { valid: true };
  },

  // Required field
  required: (value, fieldName = 'This field') => {
    if (!value || (typeof value === 'string' && !value.trim())) {
      return { valid: false, error: `${fieldName} is required` };
    }
    return { valid: true };
  },

  // Min length
  minLength: (value, length, fieldName = 'This field') => {
    if (!value || value.toString().trim().length < length) {
      return { valid: false, error: `${fieldName} must be at least ${length} characters` };
    }
    return { valid: true };
  },

  // Max length
  maxLength: (value, length, fieldName = 'This field') => {
    if (value && value.toString().trim().length > length) {
      return { valid: false, error: `${fieldName} must not exceed ${length} characters` };
    }
    return { valid: true };
  },

  // Number validation
  number: (value, fieldName = 'This field') => {
    if (value === '' || value === null || value === undefined) return { valid: true };
    if (isNaN(value)) return { valid: false, error: `${fieldName} must be a number` };
    return { valid: true };
  },

  // URL validation
  url: (url) => {
    if (!url) return { valid: false, error: 'URL is required' };
    try {
      new URL(url);
      return { valid: true };
    } catch {
      return { valid: false, error: 'Invalid URL format' };
    }
  },

  // Date validation (not in future)
  date: (date) => {
    if (!date) return { valid: false, error: 'Date is required' };
    const selectedDate = new Date(date);
    const today = new Date();
    today.setHours(0, 0, 0, 0);
    if (selectedDate > today) {
      return { valid: false, error: 'Date cannot be in the future' };
    }
    return { valid: true };
  },

  // Grade/Score validation (0-100)
  grade: (value, fieldName = 'Grade') => {
    const numVal = Number(value);
    if (isNaN(numVal)) return { valid: false, error: `${fieldName} must be a number` };
    if (numVal < 0 || numVal > 100) return { valid: false, error: `${fieldName} must be between 0 and 100` };
    return { valid: true };
  },

  // Student ID validation
  studentId: (id) => {
    if (!id) return { valid: false, error: 'Student ID is required' };
    if (!/^[A-Z0-9]{3,10}$/.test(id)) {
      return { valid: false, error: 'Student ID must be 3-10 alphanumeric characters' };
    }
    return { valid: true };
  },
};

/**
 * Form validation helper that validates multiple fields
 */
export function validateForm(formData, rules) {
  const errors = {};
  let isValid = true;

  for (const [field, rule] of Object.entries(rules)) {
    const value = formData[field];
    const validation = rule(value);
    
    if (!validation.valid) {
      errors[field] = validation.error;
      isValid = false;
    }
  }

  return { isValid, errors };
}

/**
 * Sanitize user input to prevent XSS
 */
export function sanitizeInput(input) {
  if (typeof input !== 'string') return input;
  return input
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#x27;')
    .replace(/\//g, '&#x2F;');
}

/**
 * Format phone number for display
 */
export function formatPhone(phone) {
  const cleaned = phone.replace(/\D/g, '');
  if (cleaned.length === 10) {
    return cleaned.replace(/(\d{3})(\d{3})(\d{4})/, '($1) $2-$3');
  } else if (cleaned.length === 11) {
    return cleaned.replace(/(\d{1})(\d{3})(\d{3})(\d{4})/, '+$1 ($2) $3-$4');
  }
  return phone;
}

/**
 * Parse and validate API error response
 */
export function parseApiError(error) {
  if (!error) return 'An unexpected error occurred';
  
  // Network error
  if (error.message === 'Failed to fetch' || error.message === 'Network request failed') {
    return 'Network error - please check your connection';
  }

  // API error response
  if (error.body && error.body.error) {
    if (typeof error.body.error === 'string') {
      return error.body.error;
    }
    if (error.body.error.message) {
      return error.body.error.message;
    }
  }

  if (error.body && error.body.message) {
    return error.body.message;
  }

  // HTTP status codes
  const statusErrors = {
    400: 'Bad request - please check your input',
    401: 'Unauthorized - please log in again',
    403: 'Forbidden - you do not have permission',
    404: 'Resource not found',
    409: 'Conflict - this resource already exists',
    429: 'Too many requests - please try again later',
    500: 'Server error - please try again later',
  };

  if (error.status && statusErrors[error.status]) {
    return statusErrors[error.status];
  }

  return error.message || 'An unexpected error occurred';
}
