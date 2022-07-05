import email
from operator import mod
from unittest import defaultTestLoader
from django.db import models
from django.contrib.auth.models import AbstractBaseUser


class SchoolClasses(models.Model):
    name = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class FeeType(models.Model):
    name = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class Currencies(models.Model):
    name = models.CharField(max_length=100, blank=True, default='')
    code = models.CharField(max_length=10, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class User(AbstractBaseUser):
    name = models.CharField(max_length=100, blank=True, default='')
    email = models.CharField(max_length=100, blank=True, default='')
    phone_number = models.CharField(max_length=100, blank=True, default='')
    logo = models.ImageField(upload_to='uploads/', blank=True)
    username = models.CharField(max_length=50, blank=True, default='')
    is_parent = models.BooleanField(blank=True, default=False)
    is_admin = models.BooleanField(
        blank=True, default=False)  # Use as school  id
    is_active = models.BooleanField(default=False)
    confirmed = models.BooleanField(default=False)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)
    deleted = models.DateTimeField(auto_now_add=True)


class UserPhoneNumber(models.Model):
    user_id = models.ForeignKey(
        User, on_delete=models.CASCADE, related_name="phone_numbers",)
    phone_number = models.CharField(max_length=50, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class UserEmail(models.Model):
    user_id = models.ForeignKey(
        User, on_delete=models.CASCADE, related_name="emails",)
    email = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class UserAddress(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    country = models.CharField(max_length=50, blank=True, default='')
    state = models.CharField(max_length=50, blank=True, default='')
    city = models.CharField(max_length=50, blank=True, default='')
    street = models.CharField(max_length=50, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class UserProfile(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    picture = models.ImageField(upload_to='uploads/', blank=True)
    sex = models.CharField(max_length=10, blank=True, default='')
    birth_date = models.DateField(blank=True)
    is_teacher = models.BooleanField(blank=True, default=False)
    is_secretary = models.BooleanField(blank=True, default=False)
    is_accounting = models.BooleanField(blank=True, default=False)
    is_driver = models.BooleanField(blank=True, default=False)
    is_cashier = models.BooleanField(blank=True, default=False)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)

# Modules section


class Modules(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    name = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


# Default
'''
Transport,
Fee
'''
# Payment Categories


class PaymentCategories(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    name = models.CharField(max_length=100, blank=True, default='')
    amount = models.FloatField(blank=True, default=0.0)
    currency = models.ForeignKey(Currencies, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)
    deleted = models.DateTimeField(auto_now_add=True)

# To be filled by dean of students


class ModuleToClass(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    class_id = models.ForeignKey(SchoolClasses, on_delete=models.CASCADE)
    module_id = models.ForeignKey(Modules, on_delete=models.CASCADE)
    grades = models.FloatField(blank=True)
    hours = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)
    deleted = models.DateTimeField(auto_now_add=True)

# Teacher section


class TeacherClass(models.Model):
    teach_id = models.ForeignKey(User, on_delete=models.CASCADE)
    class_id = models.ForeignKey(SchoolClasses, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)

# To be managed by  dean of students


class TeacherClassModule(models.Model):
    teach_class_id = models.ForeignKey(TeacherClass, on_delete=models.CASCADE)
    module_id = models.ForeignKey(ModuleToClass, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)

# Students  sections
# To be used by secretary account


class Students(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    first_name = models.CharField(max_length=100, blank=True, default='')
    last_name = models.CharField(max_length=100, blank=True, default='')
    birth_date = models.DateField(blank=True)
    gender = models.CharField(max_length=10, blank=True, default='')
    birth_country = models.CharField(max_length=100, blank=True, default='')
    birth_city_or_village = models.CharField(
        max_length=100, blank=True, default='')
    is_current_student = models.BooleanField(blank=True, default=False)
    left_at = models.DateTimeField()
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class StudentsAddress(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    country = models.CharField(max_length=100, blank=True, default='')
    state = models.CharField(max_length=100, blank=True, default='')
    city = models.CharField(max_length=100, blank=True, default='')
    street = models.CharField(max_length=100, blank=True, default='')
    building = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class StudentRelative(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    father_fullname = models.CharField(max_length=100, blank=True, default='')
    mother_fullname = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class StudentLevel(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    current_level = models.BooleanField(blank=True, default=False)
    studding_range_year = models.CharField(
        max_length=20, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)

# To be used by cashier


class StudentPayment(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    payment_category = models.ForeignKey(
        PaymentCategories, on_delete=models.CASCADE)
    fee_type = models.ForeignKey(FeeType, on_delete=models.CASCADE)
    description = models.CharField(max_length=500, default='', blank=True)
    amount_paid = models.FloatField()
    currency = models.ForeignKey(Currencies, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


# Default grade
'''
Fail,
Pass
'''

# To be used by a  teacher


class StudentMarksEvaluation(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    teach_id = models.ForeignKey(User, on_delete=models.CASCADE)
    module_id = models.ForeignKey(ModuleToClass, on_delete=models.CASCADE)
    student_score = models.FloatField(blank=True)
    student_grade = models.FloatField(blank=True)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


# To be used Accounting
# Default
'''
    Salary,
    Surplus
'''


class EmployeeFinancial(models.Model):
    employee_id = models.ForeignKey(User, on_delete=models.CASCADE)
    amount = models.FloatField(blank=True, default='')
    type = models.CharField(max_length=100, blank=True, default='')
    currency = models.ForeignKey(Currencies, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


# Default
'''
    Advance salary
    salary
    Credit,
    bonus
'''
'''
Month Example
JANUARY
'''


class EmployeePaymentHistory(models.Model):
    employee_id = models.ForeignKey(User, on_delete=models.CASCADE)
    amount = models.FloatField(blank=True, default='')
    month = models.CharField(max_length=100, default='', blank=True)
    year = models.CharField(max_length=100, default='', blank=True)
    payment_type = models.CharField(max_length=100, blank=True, default='')
    description = models.CharField(max_length=100, default='', blank=True)
    currency = models.ForeignKey(Currencies, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class StudentBusSubscription(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)

# To be used by cashier


class StudentBusAttendance(models.Model):
    student_subscription_id = models.ForeignKey(
        StudentBusSubscription, on_delete=models.CASCADE)
    pick_student = models.BooleanField(default=False)
    arrive_school = models.BooleanField(default=False)
    picked_at = models.DateTimeField()
    arrived_at = models.DateTimeField()


# Default
'''
student_payment,
other,
employee_payment,
bus_payment
'''


class CashReport(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    description = models.CharField(max_length=100, default='', blank=True)
    debit = models.FloatField(default=0)
    credits = models.FloatField(default=0)
    type = models.CharField(max_length=100, default='', blank=True)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


# Parent profile

class ParentProfile(models.Model):
    parent_id = models.ForeignKey(User, on_delete=models.CASCADE)
    picture = models.ImageField(upload_to='uploads/', blank=True)
    sex = models.CharField(max_length=10, blank=True, default='')
    birth_date = models.DateField(blank=True)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class ParentSchool(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    parent_profile_id = models.ForeignKey(
        ParentProfile, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class ParentSchoolStudent(models.Model):
    parent_id = models.ForeignKey(User, on_delete=models.CASCADE)
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class TeacherPost(models.Model):
    parent_id = models.ForeignKey(UserProfile, on_delete=models.CASCADE)
    teacher_id = models.ForeignKey(User, on_delete=models.CASCADE)
    post = models.TextField(default='', blank=True)
    sent_at = models.DateTimeField(auto_now_add=True)
    removed_at = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class PostCommunication(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    post = models.TextField(default='', blank=True)
    sent_at = models.DateTimeField(auto_now_add=True)
    removed_at = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)
