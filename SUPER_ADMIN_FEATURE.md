# Super Admin Feature Documentation

## Overview
The School Management System now includes a **Super Admin** feature that provides enhanced security and role-based protection for critical system administrators.

## Features Enabled

### 1. **Super Admin Role Creation**
- A new `is_super_admin` boolean field has been added to the User model
- The first admin created during system initialization is automatically set as a Super Admin
- Only Super Admins can manage other Super Admin users

### 2. **Super Admin Protection**
- **Deletion Prevention**: Super Admin users CANNOT be deleted by any other admin user
- The system will return a `403 Forbidden` error when attempting to delete a Super Admin
- This prevents accidental or malicious removal of critical system administrators

### 3. **Super Admin Management Endpoints**

#### Set Super Admin Status
```
PUT /api/admin/users/:id/super-admin
Authorization: Bearer <token>
Content-Type: application/json

Request Body:
{
  "is_super_admin": true/false
}

Response (200 OK):
{
  "message": "Super admin status updated",
  "user": {
    "id": 1,
    "email": "admin@school.com",
    "is_super_admin": true
  }
}
```

**Requirements:**
- Only Super Admin users can execute this endpoint
- Target user must have admin role to be set as Super Admin
- Returns `400 Bad Request` if attempting to set non-admin user as Super Admin
- Returns `403 Forbidden` if the requesting user is not a Super Admin

#### Get All Super Admins
```
GET /api/admin/super-admins
Authorization: Bearer <token>

Response (200 OK):
{
  "data": [
    {
      "id": 1,
      "first_name": "Admin",
      "last_name": "User",
      "email": "admin@school.com",
      "role": "admin",
      "is_super_admin": true,
      "is_active": true,
      "created_at": "2026-01-30T12:00:00Z"
    }
  ],
  "count": 1
}
```

### 4. **Delete User Protection**
When attempting to delete a user via the `/api/admin/users/:id` DELETE endpoint:
- The system first checks if the user is a Super Admin
- If the user is a Super Admin, deletion is blocked
- Returns `403 Forbidden` with message: "Super admin users cannot be deleted"

```
DELETE /api/admin/users/:id
Authorization: Bearer <token>

Response (403 Forbidden):
{
  "error": "Super admin users cannot be deleted"
}

Response (200 OK) - for non-super admin users:
{
  "message": "User deleted successfully"
}
```

## Implementation Details

### Database Schema
```sql
-- Added to users table
is_super_admin BOOLEAN DEFAULT false;
CREATE INDEX idx_users_is_super_admin ON users(is_super_admin);
```

### User Model Changes
```go
type User struct {
    // ... existing fields ...
    IsSuperAdmin bool `gorm:"default:false;index" json:"is_super_admin"`
}
```

### Service Layer
- `UserService.SetSuperAdmin(id uint, isSuperAdmin bool)` - Update super admin status
- `UserService.GetSuperAdmins()` - Fetch all super admin users
- `UserRepository.FindSuperAdmins()` - Database query for super admins

### Authorization Rules
1. Only Super Admin users can:
   - Promote other admins to Super Admin status
   - Demote Super Admins (except themselves)
   - View list of all Super Admins

2. Regular Admins can:
   - View the Super Admin status of users
   - Perform other admin operations (but cannot delete Super Admins)

3. Super Admin Protection:
   - Super Admins cannot be deleted
   - Super Admins cannot be deactivated (future enhancement)
   - At least one Super Admin should always exist in the system

## Initialization

### System Startup
When the application starts:
1. The database is auto-migrated with the new `is_super_admin` field
2. If no admin exists, the first admin created is automatically set as Super Admin
3. If an admin already exists, it is ensured to have Super Admin status
4. Logs confirm the Super Admin setup: "Admin user created successfully (set as Super Admin)"

### Default Admin Credentials
```
Email: admin@school.com (or ADMIN_EMAIL env var)
Password: admin123 (or ADMIN_PASSWORD env var)
Role: admin
Super Admin: true
```

## Error Handling

### Attempt to Delete Super Admin
```json
{
  "error": "Super admin users cannot be deleted"
}
Status: 403 Forbidden
```

### Attempt to Set Non-Admin as Super Admin
```json
{
  "error": "only admin users can be set as super admin"
}
Status: 400 Bad Request
```

### Insufficient Permissions
```json
{
  "error": "Only super admin can manage super admin status"
}
Status: 403 Forbidden
```

## Security Considerations

1. **Audit Trail**: All super admin status changes should be logged
2. **Rate Limiting**: Admin endpoints use rate limiting to prevent brute force
3. **Encryption**: Super admin status is stored as a boolean in the database
4. **Session Management**: Super Admin access requires valid JWT token
5. **RBAC**: Role-based access control ensures only Super Admins can manage this feature

## Testing Scenarios

### Test 1: Create Super Admin
```bash
# Set an admin user as Super Admin
curl -X PUT http://localhost:8080/api/admin/users/1/super-admin \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"is_super_admin": true}'
```

### Test 2: Prevent Super Admin Deletion
```bash
# Attempt to delete a Super Admin
curl -X DELETE http://localhost:8080/api/admin/users/1 \
  -H "Authorization: Bearer <token>"
# Expected: 403 Forbidden
```

### Test 3: List All Super Admins
```bash
# Get all Super Admins
curl -X GET http://localhost:8080/api/admin/super-admins \
  -H "Authorization: Bearer <token>"
```

## Related Features

- **User Management**: Full user CRUD with role-based access
- **Audit Logging**: All admin actions are logged (when audit middleware is enabled)
- **Role Middleware**: Enforces admin-only access to protected endpoints
- **Authentication**: JWT-based authentication for all protected endpoints

## Troubleshooting

### Super Admin Not Created on Startup
- Check that the database connection is working
- Verify that migrations ran successfully
- Check application logs for errors

### Cannot Set Super Admin Status
- Verify you're logged in as a Super Admin
- Check that the target user has admin role
- Ensure JWT token is valid

### User Still Deleted Despite Protection
- Clear browser cache and re-authenticate
- Verify the user is actually marked as super admin in database
- Check application logs for error details

## Future Enhancements

1. **Multi-Admin Confirmation**: Require multiple Super Admins to approve critical actions
2. **Super Admin Deactivation Lock**: Prevent deactivation of Super Admin accounts
3. **Audit Trail**: Log all Super Admin status changes
4. **Email Notifications**: Notify Super Admins of status changes
5. **Time-based Restrictions**: Limit Super Admin actions during certain hours

## Migration Path

For existing systems:
1. The `is_super_admin` field will be created with `default:false`
2. Existing admins will remain as regular admins
3. Admins can be manually promoted to Super Admin via API
4. Or run the admin promotion endpoint to set specific admins as Super Admin

```go
// Manual promotion example
PUT /api/admin/users/1/super-admin
{
  "is_super_admin": true
}
```

---

**Status**: ✅ Fully Implemented and Enabled
**Version**: 1.0
**Last Updated**: February 2, 2026
