from rest_framework import serializers
from rest_framework_simplejwt.serializers import TokenObtainPairSerializer
from models import *


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User


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


class UserPhoneNumberSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserPhoneNumber


class UserEmailSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserEmail


class UserAddressSerializer(serializers.ModelSerializer):
    class Meta:
        model = UserAddress


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
