/**
 * Reusable Form Component with Validation
 */

import React, { useState, useCallback } from 'react';
import { validateForm, parseApiError, sanitizeInput } from '../utils/validation';
import { Toast } from '../utils/toast';

export function useForm(initialValues, onSubmit, validationRules) {
  const [values, setValues] = useState(initialValues);
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState(null);

  const handleChange = useCallback((e) => {
    const { name, value, type, checked } = e.target;
    const finalValue = type === 'checkbox' ? checked : sanitizeInput(value);

    setValues(prev => ({
      ...prev,
      [name]: finalValue,
    }));

    // Clear error when field is corrected
    if (errors[name]) {
      const validation = validationRules[name]?.(finalValue);
      if (validation?.valid) {
        setErrors(prev => ({
          ...prev,
          [name]: undefined,
        }));
      }
    }
  }, [errors, validationRules]);

  const handleBlur = useCallback((e) => {
    const { name } = e.target;
    setTouched(prev => ({
      ...prev,
      [name]: true,
    }));

    // Validate field on blur
    if (validationRules[name]) {
      const validation = validationRules[name](values[name]);
      if (!validation.valid) {
        setErrors(prev => ({
          ...prev,
          [name]: validation.error,
        }));
      }
    }
  }, [values, validationRules]);

  const validate = useCallback(() => {
    const { isValid, errors: validationErrors } = validateForm(values, validationRules);
    setErrors(validationErrors);
    return isValid;
  }, [values, validationRules]);

  const handleSubmit = useCallback(async (e) => {
    if (e && e.preventDefault) {
      e.preventDefault();
    }

    // Mark all fields as touched
    const allTouched = Object.keys(initialValues).reduce((acc, key) => {
      acc[key] = true;
      return acc;
    }, {});
    setTouched(allTouched);

    // Validate form
    if (!validate()) {
      setSubmitError('Please fix the errors above');
      return;
    }

    setSubmitError(null);
    setIsSubmitting(true);

    try {
      await onSubmit(values);
    } catch (error) {
      const message = parseApiError(error);
      setSubmitError(message);
      Toast.error(message);
    } finally {
      setIsSubmitting(false);
    }
  }, [values, validate, onSubmit, initialValues]);

  const reset = useCallback(() => {
    setValues(initialValues);
    setErrors({});
    setTouched({});
    setSubmitError(null);
  }, [initialValues]);

  return {
    values,
    errors,
    touched,
    isSubmitting,
    submitError,
    handleChange,
    handleBlur,
    handleSubmit,
    setValues,
    reset,
  };
}

/**
 * FormField component for consistent form inputs
 */
export function FormField({
  label,
  name,
  type = 'text',
  error,
  touched,
  value,
  onChange,
  onBlur,
  placeholder,
  required = false,
  disabled = false,
  options = [],
  rows,
  maxLength,
  pattern,
  ...props
}) {
  const showError = touched && error;

  const baseInputStyle = {
    width: '100%',
    padding: '8px 12px',
    borderRadius: '4px',
    border: showError ? '1px solid #ef4444' : '1px solid #d1d5db',
    fontSize: '14px',
    fontFamily: 'inherit',
    boxSizing: 'border-box',
  };

  const containerStyle = {
    marginBottom: '16px',
  };

  const labelStyle = {
    display: 'block',
    marginBottom: '4px',
    fontSize: '14px',
    fontWeight: '500',
    color: '#1f2937',
  };

  const errorStyle = {
    marginTop: '4px',
    fontSize: '12px',
    color: '#ef4444',
  };

  return (
    <div style={containerStyle}>
      <label style={labelStyle}>
        {label}
        {required && <span style={{ color: '#ef4444' }}> *</span>}
      </label>

      {type === 'textarea' ? (
        <textarea
          name={name}
          value={value}
          onChange={onChange}
          onBlur={onBlur}
          placeholder={placeholder}
          disabled={disabled}
          maxLength={maxLength}
          rows={rows || 4}
          style={baseInputStyle}
          {...props}
        />
      ) : type === 'select' ? (
        <select
          name={name}
          value={value}
          onChange={onChange}
          onBlur={onBlur}
          disabled={disabled}
          style={baseInputStyle}
          {...props}
        >
          <option value="">{placeholder || 'Select option...'}</option>
          {options.map(opt => (
            <option key={opt.value} value={opt.value}>
              {opt.label}
            </option>
          ))}
        </select>
      ) : type === 'checkbox' ? (
        <label style={{ display: 'flex', alignItems: 'center', cursor: 'pointer' }}>
          <input
            type="checkbox"
            name={name}
            checked={value}
            onChange={onChange}
            onBlur={onBlur}
            disabled={disabled}
            style={{ marginRight: '8px', cursor: 'pointer' }}
            {...props}
          />
          <span style={{ fontSize: '14px' }}>{label}</span>
        </label>
      ) : (
        <input
          type={type}
          name={name}
          value={value}
          onChange={onChange}
          onBlur={onBlur}
          placeholder={placeholder}
          required={required}
          disabled={disabled}
          maxLength={maxLength}
          pattern={pattern}
          style={baseInputStyle}
          {...props}
        />
      )}

      {showError && (
        <div style={errorStyle}>
          {error}
        </div>
      )}
    </div>
  );
}

/**
 * Form component wrapper
 */
export function Form({ children, onSubmit, style, className, ...props }) {
  return (
    <form
      onSubmit={onSubmit}
      style={style}
      className={className}
      {...props}
    >
      {children}
    </form>
  );
}

/**
 * Form submit button component
 */
export function FormButton({
  type = 'submit',
  children,
  loading = false,
  disabled = false,
  variant = 'primary',
  fullWidth = true,
  ...props
}) {
  const variants = {
    primary: {
      bg: '#3b82f6',
      bgHover: '#2563eb',
      text: 'white',
    },
    danger: {
      bg: '#ef4444',
      bgHover: '#dc2626',
      text: 'white',
    },
    success: {
      bg: '#10b981',
      bgHover: '#059669',
      text: 'white',
    },
  };

  const style = variants[variant] || variants.primary;

  return (
    <button
      type={type}
      disabled={disabled || loading}
      style={{
        width: fullWidth ? '100%' : 'auto',
        padding: '10px 20px',
        backgroundColor: disabled || loading ? '#9ca3af' : style.bg,
        color: style.text,
        border: 'none',
        borderRadius: '4px',
        fontSize: '14px',
        fontWeight: '500',
        cursor: disabled || loading ? 'not-allowed' : 'pointer',
        transition: 'background-color 0.2s',
        opacity: loading ? 0.7 : 1,
      }}
      onMouseEnter={(e) => {
        if (!disabled && !loading) {
          e.target.style.backgroundColor = style.bgHover;
        }
      }}
      onMouseLeave={(e) => {
        e.target.style.backgroundColor = style.bg;
      }}
      {...props}
    >
      {loading ? '⟳ Loading...' : children}
    </button>
  );
}
