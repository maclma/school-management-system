/**
 * Enhanced API client with error handling, retry logic, and request tracking
 */

import { parseApiError } from './validation.js';
import { Toast } from './toast.js';

class APIClient {
  constructor(baseURL = '/api') {
    this.baseURL = baseURL;
    this.timeout = 30000;
    this.retryAttempts = 3;
    this.retryDelay = 1000;
    this.requestInterceptors = [];
    this.responseInterceptors = [];
  }

  /**
   * Set custom timeout
   */
  setTimeout(ms) {
    this.timeout = ms;
    return this;
  }

  /**
   * Add request interceptor
   */
  addRequestInterceptor(fn) {
    this.requestInterceptors.push(fn);
    return this;
  }

  /**
   * Add response interceptor
   */
  addResponseInterceptor(fn) {
    this.responseInterceptors.push(fn);
    return this;
  }

  /**
   * Make API request with retry logic
   */
  async request(path, options = {}) {
    const url = `${this.baseURL}${path}`;
    let lastError = null;
    let attempt = 0;

    while (attempt < this.retryAttempts) {
      try {
        return await this._makeRequest(url, options);
      } catch (error) {
        lastError = error;
        attempt++;

        // Don't retry on client errors (4xx)
        if (error.status && error.status >= 400 && error.status < 500) {
          throw error;
        }

        // Retry with exponential backoff
        if (attempt < this.retryAttempts) {
          const delay = this.retryDelay * Math.pow(2, attempt - 1);
          await new Promise(resolve => setTimeout(resolve, delay));
        }
      }
    }

    throw lastError;
  }

  /**
   * Internal method to make actual request
   */
  async _makeRequest(url, options) {
    const config = {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
    };

    // Apply request interceptors
    for (const interceptor of this.requestInterceptors) {
      interceptor(config);
    }

    // Add auth token if available
    const token = localStorage.getItem('sms_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    // Add timeout
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), this.timeout);
    config.signal = controller.signal;

    try {
      const response = await fetch(url, config);
      clearTimeout(timeoutId);

      let data = null;
      const contentType = response.headers.get('content-type');
      if (contentType?.includes('application/json')) {
        data = await response.json();
      }

      if (!response.ok) {
        const error = new Error(parseApiError({ body: data, status: response.status }));
        error.status = response.status;
        error.body = data;
        throw error;
      }

      // Apply response interceptors
      for (const interceptor of this.responseInterceptors) {
        data = interceptor(data) || data;
      }

      return data;
    } catch (error) {
      clearTimeout(timeoutId);

      // Handle abort/timeout
      if (error.name === 'AbortError') {
        const timeoutError = new Error('Request timeout - please try again');
        timeoutError.status = 0;
        throw timeoutError;
      }

      throw error;
    }
  }

  /**
   * GET request
   */
  get(path, options = {}) {
    return this.request(path, { ...options, method: 'GET' });
  }

  /**
   * POST request
   */
  post(path, data = {}, options = {}) {
    return this.request(path, {
      ...options,
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  /**
   * PUT request
   */
  put(path, data = {}, options = {}) {
    return this.request(path, {
      ...options,
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  /**
   * PATCH request
   */
  patch(path, data = {}, options = {}) {
    return this.request(path, {
      ...options,
      method: 'PATCH',
      body: JSON.stringify(data),
    });
  }

  /**
   * DELETE request
   */
  delete(path, options = {}) {
    return this.request(path, { ...options, method: 'DELETE' });
  }
}

// Create singleton instance
const apiClient = new APIClient(
  typeof import.meta !== 'undefined' && import.meta.env?.VITE_API_BASE 
    ? import.meta.env.VITE_API_BASE 
    : '/api'
);

// Add auth error interceptor
apiClient.addResponseInterceptor((data) => {
  if (data && typeof data === 'object' && data.error && data.error.code === 'UNAUTHORIZED') {
    localStorage.removeItem('sms_token');
    localStorage.removeItem('user_role');
    window.location.href = '/login';
  }
  return data;
});

export default apiClient;

/**
 * Create API endpoints object
 */
export const createAPI = () => ({
  // Authentication
  login: (email, password) =>
    apiClient.post('/auth/login', { email, password }),
  register: (data) =>
    apiClient.post('/auth/register', data),

  // Profile
  getProfile: () => apiClient.get('/profile'),
  updateProfile: (data) => apiClient.put('/profile', data),

  // Courses
  getCourses: (page = 1, limit = 50) =>
    apiClient.get(`/courses?page=${page}&limit=${limit}`),
  getCourse: (id) => apiClient.get(`/courses/${id}`),
  createCourse: (data) => apiClient.post('/courses', data),
  updateCourse: (id, data) => apiClient.put(`/courses/${id}`, data),
  deleteCourse: (id) => apiClient.delete(`/courses/${id}`),

  // Students
  getStudents: (page = 1, limit = 50) =>
    apiClient.get(`/students?page=${page}&limit=${limit}`),
  getStudent: (id) => apiClient.get(`/students/${id}`),
  createStudent: (data) => apiClient.post('/students', data),
  updateStudent: (id, data) => apiClient.put(`/students/${id}`, data),
  deleteStudent: (id) => apiClient.delete(`/students/${id}`),

  // Enrollments
  enrollStudent: (data) => apiClient.post('/enrollments', data),
  getEnrollment: (id) => apiClient.get(`/enrollments/${id}`),
  getStudentEnrollments: (studentId) =>
    apiClient.get(`/enrollments/by-student/${studentId}`),
  getCourseEnrollments: (courseId) =>
    apiClient.get(`/enrollments/by-course/${courseId}`),

  // Admin
  getAdminStats: () => apiClient.get('/admin/dashboard'),
  getAdminUsers: () => apiClient.get('/admin/users'),
  getAdminEnrollments: () => apiClient.get('/admin/enrollments'),
  approveEnrollment: (id) => apiClient.post(`/admin/enrollments/${id}/approve`),
  rejectEnrollment: (id) => apiClient.post(`/admin/enrollments/${id}/reject`),
  deleteUser: (id) => apiClient.delete(`/admin/users/${id}`),

  // Grades
  recordGrade: (data) => apiClient.post('/teacher/grades', data),
  getStudentGrades: (studentId) =>
    apiClient.get(`/grades/by-student/${studentId}`),
  getMyGrades: () => apiClient.get('/student/grades'),

  // Attendance
  recordAttendance: (data) => apiClient.post('/teacher/attendance', data),
  getStudentAttendance: (studentId) =>
    apiClient.get(`/attendance/by-student/${studentId}`),
  getMyAttendance: () => apiClient.get('/student/attendance'),

  // Assignments
  createAssignment: (data) => apiClient.post('/assignments', data),
  getAssignment: (id) => apiClient.get(`/assignments/${id}`),
  getAssignmentsByCourse: (courseId) =>
    apiClient.get(`/assignments/course/${courseId}`),
  submitAssignment: (data) => apiClient.post('/assignments/submit', data),
  getSubmissionsByStudent: () => apiClient.get('/student/assignments/submissions'),
});
