from ..serializers import NGTokenObtainPairSerializer
from rest_framework.permissions import AllowAny
from rest_framework_simplejwt.views import TokenObtainPairView


class NGObtainTokenPairView(TokenObtainPairView):
    permission_classes = (AllowAny)
    serializer_class = NGTokenObtainPairSerializer
