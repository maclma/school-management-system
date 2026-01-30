-- Advanced Features Seed Data
-- This script seeds test data for all 11 new features

-- System Settings
INSERT INTO system_settings (key, value, description, created_at, updated_at) VALUES
('app.name', 'School Management System', 'Application name', datetime('now'), datetime('now')),
('app.version', '5.0', 'Application version', datetime('now'), datetime('now')),
('email.from', 'noreply@schoolms.com', 'Default email sender', datetime('now'), datetime('now')),
('email.enabled', 'true', 'Enable email notifications', datetime('now'), datetime('now')),
('sms.enabled', 'false', 'Enable SMS notifications', datetime('now'), datetime('now')),
('backup.auto.enabled', 'true', 'Enable automatic backups', datetime('now'), datetime('now')),
('backup.schedule', 'daily', 'Backup schedule (daily/weekly/monthly)', datetime('now'), datetime('now')),
('payment.currency', 'USD', 'Currency for payments', datetime('now'), datetime('now')),
('payment.tax.rate', '0.10', 'Tax rate percentage', datetime('now'), datetime('now')),
('academic.year', '2025-2026', 'Current academic year', datetime('now'), datetime('now')),
('academic.term', 'Spring', 'Current term', datetime('now'), datetime('now'));

-- Announcements
INSERT INTO announcements (title, content, created_by, audience, priority, is_active, expires_at, created_at, updated_at) VALUES
('Welcome to New Semester', 'We are excited to announce the start of the new academic year. All students should register for their courses by the end of this week.', 1, 'all', 'high', true, 1735689600, datetime('now'), datetime('now')),
('Grade Submission Deadline', 'All teachers must submit student grades by Friday, January 31, 2026 at 5:00 PM.', 1, 'teachers', 'high', true, 1738368000, datetime('now'), datetime('now')),
('Campus Closure Notice', 'The campus will be closed on February 15, 2026 for system maintenance.', 1, 'all', 'medium', true, 1739616000, datetime('now'), datetime('now')),
('Student Financial Aid Available', 'Eligible students can now apply for financial aid. Application deadline is March 1, 2026.', 1, 'students', 'medium', true, 1741392000, datetime('now'), datetime('now')),
('Course Registration Extended', 'Registration deadline has been extended to February 7, 2026.', 1, 'students', 'high', true, 1738982400, datetime('now'), datetime('now'));

-- Messages (conversations between users)
INSERT INTO messages (sender_id, recipient_id, subject, content, is_read, created_at, updated_at) VALUES
(2, 3, 'Course Material', 'Here is the additional course material for next week. Please review before class.', true, datetime('now', '-5 days'), datetime('now', '-5 days')),
(3, 2, 'Re: Course Material', 'Thank you for sending the material. I will review it.', false, datetime('now', '-4 days'), datetime('now', '-4 days')),
(2, 4, 'Assignment 3 Clarification', 'The assignment due date has been extended to Friday. Let me know if you need any clarification.', false, datetime('now', '-2 days'), datetime('now', '-2 days')),
(4, 2, 'Re: Assignment 3 Clarification', 'Thank you for the extension. I have a question about part 2.', false, datetime('now', '-1 day'), datetime('now', '-1 day')),
(5, 1, 'New Feature Request', 'Can we add a GPA calculator to the system?', true, datetime('now', '-3 days'), datetime('now', '-3 days'));

-- Notifications
INSERT INTO notifications (user_id, title, message, type, is_read, created_at, updated_at) VALUES
(2, 'Grade Posted', 'Your grade for Advanced Calculus has been posted', 'in-app', false, datetime('now', '-1 hours'), datetime('now', '-1 hours')),
(3, 'New Message', 'You have received a new message from John (Sender 2)', 'in-app', false, datetime('now', '-2 hours'), datetime('now', '-2 hours')),
(4, 'Assignment Due Soon', 'Your assignment is due in 2 days', 'in-app', false, datetime('now', '-4 hours'), datetime('now', '-4 hours')),
(2, 'Enrollment Approved', 'Your enrollment in Physics 101 has been approved', 'in-app', true, datetime('now', '-1 days'), datetime('now', '-1 days')),
(3, 'Attendance Alert', 'You are missing attendance in Database Systems (3 absences)', 'in-app', false, datetime('now', '-3 hours'), datetime('now', '-3 hours'));

-- Payments
INSERT INTO payments (student_id, amount, status, description, due_date, created_at, updated_at) VALUES
(2, 2500.00, 'paid', 'Spring 2026 Tuition', 1735689600, datetime('now', '-10 days'), datetime('now', '-10 days')),
(2, 150.00, 'paid', 'Laboratory Fee', 1735689600, datetime('now', '-8 days'), datetime('now', '-8 days')),
(3, 2500.00, 'pending', 'Spring 2026 Tuition', 1738368000, datetime('now'), datetime('now')),
(3, 150.00, 'overdue', 'Laboratory Fee', 1735689600, datetime('now', '-5 days'), datetime('now', '-5 days')),
(4, 2500.00, 'paid', 'Spring 2026 Tuition', 1738368000, datetime('now', '-2 days'), datetime('now', '-2 days')),
(5, 2500.00, 'cancelled', 'Spring 2026 Tuition (Refunded)', 1738368000, datetime('now', '-15 days'), datetime('now', '-15 days'));

-- TimeTable (Class Schedule)
INSERT INTO timetables (course_id, teacher_id, day_of_week, start_time, end_time, classroom, location, created_at, updated_at) VALUES
(1, 2, 'Monday', '09:00', '10:30', '101', 'Building A', datetime('now'), datetime('now')),
(1, 2, 'Wednesday', '09:00', '10:30', '101', 'Building A', datetime('now'), datetime('now')),
(1, 2, 'Friday', '09:00', '10:30', '101', 'Building A', datetime('now'), datetime('now')),
(2, 3, 'Tuesday', '11:00', '12:30', '202', 'Building B', datetime('now'), datetime('now')),
(2, 3, 'Thursday', '11:00', '12:30', '202', 'Building B', datetime('now'), datetime('now')),
(3, 2, 'Monday', '14:00', '15:30', '301', 'Building C', datetime('now'), datetime('now')),
(3, 2, 'Wednesday', '14:00', '15:30', '301', 'Building C', datetime('now'), datetime('now')),
(4, 4, 'Tuesday', '13:00', '14:30', '105', 'Building A', datetime('now'), datetime('now')),
(4, 4, 'Thursday', '13:00', '14:30', '105', 'Building A', datetime('now'), datetime('now'));

-- Grade Transcripts (Academic Records)
INSERT INTO grade_transcripts (student_id, academic_year, term, gpa, total_credits, completed_credits, created_at, updated_at) VALUES
(2, '2024-2025', 'Fall', 3.75, 30, 30, datetime('now'), datetime('now')),
(2, '2025-2026', 'Spring', 3.80, 15, 0, datetime('now'), datetime('now')),
(3, '2024-2025', 'Fall', 3.45, 30, 30, datetime('now'), datetime('now')),
(3, '2025-2026', 'Spring', 3.50, 15, 0, datetime('now'), datetime('now')),
(4, '2024-2025', 'Fall', 3.90, 30, 30, datetime('now'), datetime('now')),
(4, '2025-2026', 'Spring', 3.92, 15, 0, datetime('now'), datetime('now')),
(5, '2024-2025', 'Fall', 2.95, 30, 30, datetime('now'), datetime('now')),
(5, '2025-2026', 'Spring', 3.15, 15, 0, datetime('now'), datetime('now'));

-- Backups (Database Backups)
INSERT INTO backups (backup_name, status, size_bytes, file_path, backed_up_at, created_at, updated_at) VALUES
('backup_2026_01_29_automatic', 'completed', 5242880, '/backups/2026/01/29/backup_automatic.sql', 1737986638, datetime('now', '-2 hours'), datetime('now', '-2 hours')),
('backup_2026_01_28_automatic', 'completed', 5120000, '/backups/2026/01/28/backup_automatic.sql', 1737900238, datetime('now', '-1 days'), datetime('now', '-1 days')),
('backup_2026_01_27_automatic', 'completed', 4998976, '/backups/2026/01/27/backup_automatic.sql', 1737813838, datetime('now', '-2 days'), datetime('now', '-2 days')),
('backup_2026_01_29_manual', 'in_progress', 0, '/backups/2026/01/29/backup_manual_running.sql', NULL, datetime('now', '-30 minutes'), datetime('now', '-30 minutes')),
('backup_2026_01_26_automatic', 'completed', 4856789, '/backups/2026/01/26/backup_automatic.sql', 1737727438, datetime('now', '-3 days'), datetime('now', '-3 days'));

-- Import Batches (CSV Import Tracking)
INSERT INTO import_batches (entity_type, status, total_records, successful_records, failed_records, error_log, created_at, updated_at) VALUES
('students', 'completed', 50, 50, 0, NULL, datetime('now', '-7 days'), datetime('now', '-7 days')),
('teachers', 'completed', 15, 15, 0, NULL, datetime('now', '-6 days'), datetime('now', '-6 days')),
('courses', 'completed', 20, 19, 1, 'Row 5: Invalid department ID', datetime('now', '-5 days'), datetime('now', '-5 days')),
('enrollments', 'completed', 120, 118, 2, 'Rows 45,89: Student already enrolled', datetime('now', '-4 days'), datetime('now', '-4 days')),
('grades', 'in_progress', 300, 280, 0, NULL, datetime('now'), datetime('now'));

-- Audit Logs (Compliance & Audit Trail)
INSERT INTO audit_logs (user_id, action, entity_type, entity_id, details, ip_address, created_at) VALUES
(1, 'LOGIN', 'user', 1, 'Admin login successful', '192.168.1.100', datetime('now', '-2 hours')),
(1, 'CREATE', 'announcement', 1, 'Created new announcement: Welcome to New Semester', '192.168.1.100', datetime('now', '-1 day')),
(2, 'UPDATE', 'grade', 5, 'Grade updated from 85 to 88', '192.168.1.101', datetime('now', '-3 hours')),
(3, 'VIEW', 'student_record', 2, 'Viewed student details', '192.168.1.102', datetime('now', '-1 hour')),
(1, 'DELETE', 'message', 3, 'Deleted message ID 3', '192.168.1.100', datetime('now', '-30 minutes')),
(4, 'EXPORT', 'grades', NULL, 'Exported class grades to CSV', '192.168.1.103', datetime('now', '-15 minutes')),
(1, 'APPROVE', 'enrollment', 15, 'Approved student enrollment', '192.168.1.100', datetime('now', '-45 minutes'));

-- Summary: Advanced Features Seeding Complete
-- Total records seeded:
-- - System Settings: 11
-- - Announcements: 5
-- - Messages: 5
-- - Notifications: 5
-- - Payments: 6
-- - TimeTable entries: 9
-- - Grade Transcripts: 8
-- - Backups: 5
-- - Import Batches: 5
-- - Audit Logs: 7
