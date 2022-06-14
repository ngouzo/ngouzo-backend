from django.urls import path
from .views import LOGIN

urlpatterns = [
    path('login/',  LOGIN)
]
