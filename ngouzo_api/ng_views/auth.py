from contextvars import Token
import json
from urllib import response
from django.contrib.auth.models import User
from django.forms import ValidationError
from rest_framework import generics
from rest_framework.permissions import AllowAny
from rest_framework.decorators import api_view, permission_classes
from django.contrib.auth import login
from django.contrib.auth.hashers import check_password
from django.views.decorators.csrf import csrf_exempt, csrf_protect


from ..serializers import UserSerializer


class RegisterView(generics.CreateAPIView):
    queryset = User.objects.all()
    permission_classes = [AllowAny]
    serializer_class = UserSerializer


@api_view(["POST"])
@permission_classes([AllowAny])
@csrf_exempt
def login_user(request):
    data = {}
    reqBody = json.loads(request.body)
    email1 = reqBody['Email_Address']
    print(email1)
    password = reqBody['password']
    try:
        Account = User.objects.get(Email_Address=email1)
    except BaseException as e:
        raise ValidationError({"400": f'{str(e)}'})

    token = Token.objects.get_or_create(user=Account)[0].key
    print(token)
    if not check_password(password, Account.password):
        raise ValidationError({"message": "Incorrect Login credentials"})

    if Account:
        if Account.is_active:
            print(request.user)
            login(request, Account)
            data["message"] = "user logged in"
            data["email_address"] = Account.email

            Res = {"data": data, "token": token}

            return response(Res)

        else:
            raise ValidationError({"400": f'Account not active'})

    else:
        raise ValidationError({"400": f"Account doesn't exist"})
