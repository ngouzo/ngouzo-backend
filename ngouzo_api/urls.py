from django.urls import path
from rest_framework_simplejwt.views import (
    TokenObtainPairView,
    TokenRefreshView,
)
from .views import REGISTER_VIEW, LOGIN_VIEW

urlpatterns = [
    path('api/token/', TokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('api/token/refresh/', TokenRefreshView.as_view(), name='token_refresh'),
    path('register/', REGISTER_VIEW.as_view(), name='auth_register'),
    path('login/', LOGIN_VIEW, name="auth_login"),
]
