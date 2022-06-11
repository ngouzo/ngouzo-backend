from operator import mod
from unittest import defaultTestLoader
from django.db import models
from pygments.lexers import get_all_lexers
from pygments.styles import get_all_styles

LEXERS = [item for item in get_all_lexers() if item[1]]
STYLE_CHOICES = sorted([(item, item) for item in get_all_styles()])


class SchoolClasses(models.Model):
    name = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class Currencies(models.Model):
    name = models.CharField(max_length=100, blank=True, default='')
    code = models.CharField(max_length=10, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class User(models.Model):
    name = models.CharField(max_length=100, blank=True, default='')
    logo = models.ImageField(upload_to='uploads/', blank=True)
    first_name = models.CharField(max_length=100, blank=True, default='')
    last_name = models.CharField(max_length=100, default=True, default='')
    username = models.CharField(max_length=50, blank=True, default='')
    password = models.CharField(max_length=16, blank=True, default='')
    is_admin = models.BooleanField(
        blank=True, default=False)  # Use as school  id
    is_active = models.BooleanField(default=False)
    confirmed = models.BooleanField(default=False)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)
    deleted = models.DateTimeField(auto_now_add=True)


class UserPhoneNumber(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    phone_number = models.CharField(max_length=50, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class UserEmail(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
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
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    picture = models.ImageField(upload_to='uploads/', blank=True)
    sex = models.CharField(max_length=10, blank=True, default='')
    birth_date = models.DateField(blank=True)
    is_teacher = models.BooleanField(blank=True, default=False)
    is_secretary = models.BooleanField(blank=True, default=False)
    is_accounting = models.BooleanField(blank=True, default=False)
    is_driver = models.BooleanField(blank=True, default=False)
    is_parent = models.BooleanField(blank=True, default=False)
    is_cashier = models.BooleanField(blank=True, default=False)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)

# Modules section


class Modules(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    name = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


# Payment Categories


class PaymentCategories(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    name = models.CharField(max_length=100, blank=True, default='')
    amount = models.FloatField(blank=True, default=0.0)
    currency = models.ForeignKey(Currencies, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)
    deleted = models.DateTimeField(auto_now_add=True)
# Students  sections


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


class StudentRelative(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    father_fullname = models.CharField(max_length=100, blank=True, default='')
    mother_fullname = models.CharField(max_length=100, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class StudentLevel(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    current_level = models.BooleanField(blank=True, default=False)
    studding_year = models.CharField(max_length=20, blank=True, default='')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class StudentPayment(models.Model):
    student_id = models.ForeignKey(Students, on_delete=models.CASCADE)
    payment_category = models.ForeignKey(
        PaymentCategories, on_delete=models.CASCADE)
    amount_paid = models.FloatField()
    currency = models.ForeignKey(Currencies, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class ModuleToClass(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    class_id = models.ForeignKey(SchoolClasses, on_delete=models.CASCADE)
    module_id = models.ForeignKey(Modules, on_delete=models.CASCADE)


# Teacher section
class TeacherClass(models.Model):
    school_id = models.ForeignKey(User, on_delete=models.CASCADE)
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    class_id = models.ForeignKey(SchoolClasses, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)


class TeacherClassModule(models.Model):
    teach_class_id = models.ForeignKey(TeacherClass, on_delete=models.CASCADE)
    module_id = models.ForeignKey(ModuleToClass, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now_add=True)
