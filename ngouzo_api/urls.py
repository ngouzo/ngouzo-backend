from django.urls import path
from rest_framework_simplejwt.views import (
    TokenObtainPairView,
    TokenRefreshView,
)

from .ng_views.ngObtainTokenPair import NGObtainTokenPairView
from .views import REGISTER_VIEW, LOGIN_VIEW

urlpatterns = [
    path('api/login/', NGObtainTokenPairView.as_view(), name='token_obtain_pair'),
    path('api/login/refresh/', TokenRefreshView.as_view(), name='token_refresh'),
    path('register/', REGISTER_VIEW.as_view(), name='auth_register'),

]
