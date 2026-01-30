/**
 * Error Boundary Component
 * Catches and displays errors gracefully
 */

import React from 'react';

export class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      hasError: false,
      error: null,
      errorInfo: null,
      errorCount: 0,
    };
  }

  static getDerivedStateFromError(error) {
    return { hasError: true };
  }

  componentDidCatch(error, errorInfo) {
    this.setState(prevState => ({
      error,
      errorInfo,
      errorCount: prevState.errorCount + 1,
    }));

    // Log to external service in production
    if (process.env.NODE_ENV === 'production') {
      console.error('Error logged:', error, errorInfo);
    }
  }

  resetError = () => {
    this.setState({
      hasError: false,
      error: null,
      errorInfo: null,
    });
  };

  render() {
    if (this.state.hasError) {
      return (
        <div style={{
          padding: '20px',
          margin: '20px',
          border: '1px solid #ef4444',
          borderRadius: '8px',
          backgroundColor: '#fef2f2',
          color: '#7f1d1d',
        }}>
          <h2 style={{ marginTop: 0, color: '#dc2626' }}>
            Oops! Something went wrong
          </h2>
          
          {process.env.NODE_ENV === 'development' && (
            <details style={{ whiteSpace: 'pre-wrap', marginTop: '10px' }}>
              <summary style={{ cursor: 'pointer', fontWeight: 'bold' }}>
                Error Details (Development Only)
              </summary>
              <p>{this.state.error?.toString()}</p>
              <p>{this.state.errorInfo?.componentStack}</p>
            </details>
          )}

          <button
            onClick={this.resetError}
            style={{
              marginTop: '10px',
              padding: '8px 16px',
              backgroundColor: '#ef4444',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
              fontSize: '14px',
            }}
          >
            Try again
          </button>

          {this.state.errorCount > 3 && (
            <p style={{ marginTop: '10px', fontSize: '14px' }}>
              Multiple errors detected. Please <a href="/">return home</a>.
            </p>
          )}
        </div>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
