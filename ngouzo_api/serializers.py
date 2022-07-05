import email
from django.forms import EmailField
from pkg_resources import require
from rest_framework import serializers
from rest_framework.validators import UniqueValidator
from rest_framework_simplejwt.serializers import TokenObtainPairSerializer
from django.contrib.auth.password_validation import validate_password

from .models import *


class UserEmailSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserEmail
        fields = "__all__"


class UserPhoneNumberSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserPhoneNumber
        fields = "__all__"


class UserSerializer(serializers.ModelSerializer):
    email = serializers.EmailField(
        required=True,
        validators=[UniqueValidator(queryset=User.objects.all())]
    )
    phone_number = serializers.CharField(
        required=True,
        validators=[UniqueValidator(queryset=User.objects.all())]
    )

    password = serializers.CharField(
        write_only=True, required=True, validators=[validate_password])
    password2 = serializers.CharField(write_only=True, required=True)

    class Meta:
        model = User
        fields = ('username', 'password', 'password2',
                  'name', 'phone_number', 'email')
        extra_kwargs = {
            'name': {'required': True}
        }

    def validate(self, attrs):
        if attrs['password'] != attrs['password2']:
            raise serializers.ValidationError(
                {"password": "Password fields didn't match."})
        return attrs

    def create(self, validate_data):
        user = User.objects.create(
            username=validate_data['username'],
            name=validate_data['name'],
            email=validate_data['email'],
            phone_number=validate_data['phone_number']
        )
        user.set_password(validate_data['password'])
        user.save()
        return user


class UserProfileSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserProfile


class SchoolClassesSerializer(serializers.ModelSerializer):
    class Meta:
        model = SchoolClasses


class FeeTypeSerializer(serializers.ModelSerializer):
    class Meta:
        model = FeeType


class CurrencySerializer(serializers.ModelSerializer):
    class Meta:
        model = Currencies


class UserAddressSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserAddress
        fields = "__all__"


class ModuleSerializer(serializers.ModelSerializer):
    class Meta:
        model = Modules


class PaymentCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = PaymentCategories


class ModuleToClassSerializer(serializers.ModelField):
    class Meta:
        model = ModuleToClass


class TeacherClassSerializer(serializers.ModelField):
    class Meta:
        model = TeacherClass


class TeacherClassModuleSerializer(serializers.ModelSerializer):
    class Meta:
        model = TeacherClassModule


class StudentsSerializer(serializers.ModelField):
    class Meta:
        model = Students


class StudentAddressSerializer(serializers.ModelSerializer):
    class Meta:
        model = StudentsAddress


class StudentRelativeSerializer(serializers.ModelSerializer):
    class Meta:
        model = StudentRelative


class StudentLevelSerializer(serializers.ModelSerializer):
    class Meta:
        model = StudentLevel


class StudentPaymentSerial(serializers.ModelSerializer):
    class Meta:
        model = StudentPayment


class StudentMarsEvaluationSerializer(serializers.ModelSerializer):
    class Meta:
        model = StudentMarksEvaluation


class EmployeeFinancialSerializer(serializers.ModelSerializer):
    class Meta:
        model = EmployeeFinancial


class EmployeePaymentHistorySerializer(serializers.ModelSerializer):
    class Meta:
        model = EmployeePaymentHistory


class StudentBusSubscriptionSerializer(serializers.ModelSerializer):
    class Meta:
        model = StudentBusSubscription


class StudentBusAttendanceSerializer(serializers.ModelSerializer):
    class Meta:
        model = StudentBusAttendance


class CashReportSerializer(serializers.ModelSerializer):
    class Meta:
        model = CashReport


class ParentProfileSerializer(serializers.ModelSerializer):
    class Meta:
        model = ParentProfile


class ParentSchoolSerializer(serializers.ModelSerializer):
    class Meta:
        model = ParentSchool


class ParentSchoolStudent(serializers.ModelSerializer):
    class Meta:
        model = ParentSchoolStudent


class TeacherPostSerializer(serializers.ModelSerializer):
    class Meta:
        model = TeacherPost


class PostCommunicationSerializer(serializers.ModelSerializer):
    class Meta:
        model = PostCommunication

# Create token obtain


class NGTokenObtainPairSerializer(TokenObtainPairSerializer):
    @classmethod
    def get_token(cls, user):
        token = super(NGTokenObtainPairSerializer).get_token(user)
        token['username'] = user.username
        return token
