// lib/utils/app_theme.dart
import 'package:flutter/material.dart';

class AppColors {
  // Brand Colors (from logo)
  static const neonCyan = Color(0xFF00F5FF);
  static const neonPurple = Color(0xFFBF5FFF);
  static const indigo = Color(0xFF4F46E5);
  static const hotPink = Color(0xFFFF2D78);

  // Background
  static const bg = Color(0xFF0A0A14);
  static const bg2 = Color(0xFF0E0E1E);
  static const surface = Color(0xFF141428);
  static const surfaceHover = Color(0xFF1A1A35);
  static const border = Color(0xFF252540);

  // Text
  static const textPrimary = Color(0xFFEAEAF5);
  static const textMuted = Color(0xFF6B6B8F);

  // Status
  static const success = Color(0xFF00E5A0);
  static const warning = Color(0xFFFFB800);
  static const danger = Color(0xFFFF4444);

  // Gradients
  static const gradientPrimary = LinearGradient(
    colors: [indigo, neonPurple],
    begin: Alignment.centerLeft,
    end: Alignment.centerRight,
  );

  static const gradientCyan = LinearGradient(
    colors: [neonCyan, indigo],
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
  );

  static const gradientBg = LinearGradient(
    colors: [bg, bg2],
    begin: Alignment.topCenter,
    end: Alignment.bottomCenter,
  );
}

class AppTheme {
  static ThemeData get dark => ThemeData(
        useMaterial3: true,
        brightness: Brightness.dark,
        scaffoldBackgroundColor: AppColors.bg,
        fontFamily: 'PlusJakartaSans',
        colorScheme: const ColorScheme.dark(
          primary: AppColors.indigo,
          secondary: AppColors.neonCyan,
          surface: AppColors.surface,
          background: AppColors.bg,
          error: AppColors.danger,
          onPrimary: Colors.white,
          onSecondary: AppColors.bg,
          onSurface: AppColors.textPrimary,
          onBackground: AppColors.textPrimary,
        ),
        appBarTheme: const AppBarTheme(
          backgroundColor: AppColors.bg,
          elevation: 0,
          centerTitle: false,
          titleTextStyle: TextStyle(
            fontFamily: 'Syne',
            fontSize: 20,
            fontWeight: FontWeight.w800,
            color: AppColors.textPrimary,
          ),
          iconTheme: IconThemeData(color: AppColors.textPrimary),
        ),
        bottomNavigationBarTheme: const BottomNavigationBarThemeData(
          backgroundColor: AppColors.surface,
          selectedItemColor: AppColors.neonCyan,
          unselectedItemColor: AppColors.textMuted,
          type: BottomNavigationBarType.fixed,
          elevation: 0,
        ),
        cardTheme: CardTheme(
          color: AppColors.surface,
          elevation: 0,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(20),
            side: const BorderSide(color: AppColors.border, width: 1),
          ),
        ),
        textTheme: const TextTheme(
          displayLarge: TextStyle(
            fontFamily: 'Syne',
            fontSize: 32,
            fontWeight: FontWeight.w800,
            color: AppColors.textPrimary,
          ),
          headlineMedium: TextStyle(
            fontFamily: 'Syne',
            fontSize: 22,
            fontWeight: FontWeight.w700,
            color: AppColors.textPrimary,
          ),
          titleLarge: TextStyle(
            fontSize: 18,
            fontWeight: FontWeight.w600,
            color: AppColors.textPrimary,
          ),
          bodyLarge: TextStyle(
            fontSize: 15,
            color: AppColors.textPrimary,
          ),
          bodyMedium: TextStyle(
            fontSize: 13,
            color: AppColors.textMuted,
          ),
        ),
        inputDecorationTheme: InputDecorationTheme(
          filled: true,
          fillColor: AppColors.surface,
          border: OutlineInputBorder(
            borderRadius: BorderRadius.circular(12),
            borderSide: const BorderSide(color: AppColors.border),
          ),
          enabledBorder: OutlineInputBorder(
            borderRadius: BorderRadius.circular(12),
            borderSide: const BorderSide(color: AppColors.border),
          ),
          focusedBorder: OutlineInputBorder(
            borderRadius: BorderRadius.circular(12),
            borderSide: const BorderSide(color: AppColors.indigo, width: 2),
          ),
          labelStyle: const TextStyle(color: AppColors.textMuted),
          hintStyle: const TextStyle(color: AppColors.textMuted),
        ),
        elevatedButtonTheme: ElevatedButtonThemeData(
          style: ElevatedButton.styleFrom(
            backgroundColor: AppColors.indigo,
            foregroundColor: Colors.white,
            elevation: 0,
            padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 14),
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(12),
            ),
            textStyle: const TextStyle(
              fontWeight: FontWeight.w600,
              fontSize: 15,
            ),
          ),
        ),
      );
}

class AppConstants {
  static const String baseUrl = 'https://api.ploykong.com/api/v1';
  static const String appName = 'PloyKong';
  static const Duration animDuration = Duration(milliseconds: 300);
}
