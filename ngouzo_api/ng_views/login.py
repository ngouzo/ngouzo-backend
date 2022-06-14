from django.http import HttpResponse, JsonResponse
from django.views.decorators.csrf import csrf_exempt
from rest_framework.parsers import JSONParser


@csrf_exempt
def user_login(request):
    if request.method == 'GET':
        return JsonResponse({"msg":  "Welcome to ngouzo"})
